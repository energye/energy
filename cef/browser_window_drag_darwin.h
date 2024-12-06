extern void GoLog(char* message);
extern bool CanDrag(void* nsWindow);
extern void SetCanDrag(void* nsWindow, bool);
extern int32_t GetTitlebarHeight(void* nsWindow);
extern void CheckDraggableRegions(void* nsWindow, int32_t mouseX, int32_t mouseY);