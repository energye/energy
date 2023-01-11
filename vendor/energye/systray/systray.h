#include "stdbool.h"

extern void systray_ready();
extern void systray_on_exit();
extern void systray_menu_item_selected(int menu_id);
extern void systray_on_click();
extern void systray_on_rclick();

void registerSystray(void);
void nativeEnd(void);
int nativeLoop(void);
void nativeStart(void);

void setIcon(const char* iconBytes, int length, bool template);
void setMenuItemIcon(const char* iconBytes, int length, int menuId, bool template);
void setTitle(char* title);
void setTooltip(char* tooltip);
void add_or_update_menu_item(int menuId, int parentMenuId, char* title, char* tooltip, char* shortcutKey, short disabled, short checked, short isCheckable);
void add_separator(int menuId);
void hide_menu_item(int menuId);
void show_menu_item(int menuId);
void reset_menu();
void create_menu();
void show_menu();
void set_menu_nil();
void quit();
void enable_on_click();
