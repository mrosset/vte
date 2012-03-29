#include <gtk/gtk.h>
#include <vte/vte.h>

static void
resize_window(GtkWidget *widget, guint width, guint height, gpointer data)
{
	VteTerminal *terminal;
	printf("RESIZE");
	fflush(stdout);
	if ((GTK_IS_WINDOW(data)) && (width >= 2) && (height >= 2)) {
		gint owidth, oheight, char_width, char_height, column_count, row_count;
		GtkBorder *inner_border;

		terminal = VTE_TERMINAL(widget);

		gtk_window_get_size(GTK_WINDOW(data), &owidth, &oheight);

		/* Take into account border overhead. */
		char_width = vte_terminal_get_char_width (terminal);
		char_height = vte_terminal_get_char_height (terminal);
		column_count = vte_terminal_get_column_count (terminal);
		row_count = vte_terminal_get_row_count (terminal);
		gtk_widget_style_get (widget, "inner-border", &inner_border, NULL);

		owidth -= char_width * column_count;
		oheight -= char_height * row_count;
		if (inner_border != NULL) {
			owidth -= inner_border->left + inner_border->right;
			oheight -= inner_border->top + inner_border->bottom;
		}
		gtk_window_resize(GTK_WINDOW(data),
				  width + owidth, height + oheight);
		gtk_border_free (inner_border);
	}
}

static void
destroy_and_quit(VteTerminal *terminal, GtkWidget *window)
{
	const char *output_file = g_object_get_data (G_OBJECT (terminal), "output_file");

	if (output_file) {
		GFile *file;
		GOutputStream *stream;
		GError *error = NULL;

		file = g_file_new_for_commandline_arg (output_file);
		stream = G_OUTPUT_STREAM (g_file_replace (file, NULL, FALSE, G_FILE_CREATE_NONE, NULL, &error));

		if (stream) {
			vte_terminal_write_contents (terminal, stream,
						     VTE_TERMINAL_WRITE_DEFAULT,
						     NULL, &error);
			g_object_unref (stream);
		}

		if (error) {
			g_printerr ("%s\n", error->message);
			g_error_free (error);
		}

		g_object_unref (file);
	}

	gtk_widget_destroy (window);
	gtk_main_quit ();
}

static void
child_exited(GtkWidget *terminal, gpointer window) 
{
	vte_terminal_get_child_exit_status (VTE_TERMINAL (terminal));
	destroy_and_quit(VTE_TERMINAL (terminal), GTK_WIDGET (window));
}

int
main(int argc, char **argv) 
{
	VteTerminal *terminal;
	GtkWidget *widget, *window, *scrolled_window;
	GdkScreen *screen;
	GdkColormap *colormap;

	VtePtyFlags pty_flags = VTE_PTY_DEFAULT;
	GtkPolicyType scrollbar_policy = GTK_POLICY_ALWAYS;

	char **command_argv = NULL;
	int command_argc;
	const char *command = NULL;

	GPid pid = -1;
	GError *err = NULL;

	gtk_init (&argc, &argv);

	widget = vte_terminal_new();
	window = gtk_window_new(GTK_WINDOW_TOPLEVEL);

	gtk_container_set_resize_mode(GTK_CONTAINER(window),
			GTK_RESIZE_IMMEDIATE);
	screen = gtk_widget_get_screen (window);
	colormap = gdk_screen_get_rgba_colormap (screen);


	scrolled_window = gtk_scrolled_window_new (NULL, NULL);
	gtk_scrolled_window_set_policy(GTK_SCROLLED_WINDOW(scrolled_window),
			GTK_POLICY_NEVER, scrollbar_policy);
	gtk_container_add(GTK_CONTAINER(window), GTK_WIDGET(scrolled_window));

	terminal = VTE_TERMINAL (widget);

	command = vte_get_user_shell();
	g_shell_parse_argv(command, &command_argc, &command_argv, &err);
	vte_terminal_fork_command_full(terminal,
			pty_flags,
			NULL,
			command_argv,
			NULL,
			G_SPAWN_SEARCH_PATH,
			NULL, NULL,
			&pid,
			&err);

	//g_signal_connect(widget, "resize-window",G_CALLBACK(resize_window), window);
	//g_signal_connect(widget, "child-exited", G_CALLBACK(child_exited), window);

	gtk_container_add(GTK_CONTAINER(scrolled_window), GTK_WIDGET(terminal));

	gtk_widget_show_all(window);

	gtk_main();
}
