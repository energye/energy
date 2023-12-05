#import "WindowController.h"

@interface WindowController () <NSTouchBarDelegate>
@property (copy) void (^handler)(char *);
@property (retain) NSDictionary* goData;
@property (copy) NSRegularExpression* hexRegex;
@property (copy) NSDictionary* identifierMapping;
@property (copy) NSDictionary* imageMapping;
@property (retain) NSLock *lock;
@end

@implementation WindowController

static NSTouchBarItemIdentifier standardOtherItems = @"com.energy.touchbar.other_items";
static NSTouchBarItemIdentifier standardSpaceSmall = @"com.energy.touchbar.small_space";
static NSTouchBarItemIdentifier standardSpaceLarge = @"com.energy.touchbar.large_space";
static NSTouchBarItemIdentifier standardSpaceFlexible = @"com.energy.touchbar.flexible_space";
static NSTouchBarItemIdentifier standardCandidateList = @"com.energy.touchbar.candidates";
static NSTouchBarItemIdentifier standardCharacterPicker = @"com.energy.touchbar.char_picker";
static NSTouchBarItemIdentifier standardTextFormat = @"com.energy.touchbar.text_format";
static NSTouchBarItemIdentifier standardTextAlignment = @"com.energy.touchbar.text_align";
static NSTouchBarItemIdentifier standardTextColorPicker = @"com.energy.touchbar.text_color";
static NSTouchBarItemIdentifier standardTextList = @"com.energy.touchbar.text_list";
static NSTouchBarItemIdentifier standardTextStyle = @"com.energy.touchbar.text_style";

static NSTouchBarItemIdentifier prefixWidget = @"com.energy.touchbar";
static NSTouchBarItemIdentifier prefixButton = @"com.energy.touchbar.button.";
static NSTouchBarItemIdentifier prefixCandidates = @"com.energy.touchbar.candidates.";
static NSTouchBarItemIdentifier prefixColorpicker = @"com.energy.touchbar.colorpicker.";
static NSTouchBarItemIdentifier prefixCustom = @"com.energy.touchbar.custom.";
static NSTouchBarItemIdentifier prefixGroup = @"com.energy.touchbar.group.";
static NSTouchBarItemIdentifier prefixLabel = @"com.energy.touchbar.label.";
static NSTouchBarItemIdentifier prefixPicker = @"com.energy.touchbar.picker.";
static NSTouchBarItemIdentifier prefixPopover = @"com.energy.touchbar.popover.";
static NSTouchBarItemIdentifier prefixScrubber = @"com.energy.touchbar.scrubber.";
static NSTouchBarItemIdentifier prefixSegmented = @"com.energy.touchbar.segmented.";
static NSTouchBarItemIdentifier prefixSharer = @"com.energy.touchbar.sharer.";
static NSTouchBarItemIdentifier prefixSlider = @"com.energy.touchbar.slider.";
static NSTouchBarItemIdentifier prefixStepper = @"com.energy.touchbar.stepper.";

- (id)initWithData:(const char *)data andHandler:(void (^)(char *))handler error:(NSError**)error {
  if ((self = [WindowController alloc]) == nil) {
    return nil;
  }
  self.lock = [[NSLock alloc] init];
  self.handler = handler;
  NSError* err = [self setData:data];
  if (err != nil) {
    if (error != nil) {
      *error = err;
    }
    return nil;
  }
  err = [self initConst];
  if (err != nil) {
    if (error != nil) {
      *error = err;
    }
    return nil;
  }
  [self initMapping];
  return self;
}

- (NSError*)update:(NSTouchBar*)touchBar withData:(const char *)data {
  [self.lock lock];
  NSError* err = [self setData:data];
  if (err == nil) {
    [self setupTouchBar:touchBar];
  }
  [self.lock unlock];
  if (err == nil) {
    [self updateItems:touchBar];
  }
  return err;
}

- (NSArray<NSTouchBarItemIdentifier>*)mapIdentifiers:(NSArray*)input {
  NSMutableArray* output = [[NSMutableArray alloc] init];
  [output setArray:input];
  for (int i = 0; i < [output count]; ++i) {
    NSTouchBarItemIdentifier newIdentifier = [self transformIdentifier:[output objectAtIndex:i]];
    if (newIdentifier != nil) {
      [output replaceObjectAtIndex:i withObject:newIdentifier];
    }
  }
  return output;
}

- (void)setupTouchBar:(NSTouchBar*)touchBar {
  [touchBar setDefaultItemIdentifiers:[self mapIdentifiers:[self.goData objectForKey:@"Default"]]];
  [touchBar setPrincipalItemIdentifier:[self transformIdentifier:[self.goData objectForKey:@"Principal"]]];
  [touchBar setEscapeKeyReplacementItemIdentifier:[self transformIdentifier:[self.goData objectForKey:@"Escape"]]];
}

- (NSError*)setData:(const char *)rawData {
  NSData* data = [NSData dataWithBytes:(const void *)rawData length:sizeof(unsigned char) * strlen(rawData)];

  NSError* error = nil;
  self.goData = [NSJSONSerialization JSONObjectWithData:data options:0 error:&error];
  if (error != nil) {
    return error;
  }
  return nil;
}

- (NSTouchBar*)makeTouchBar {
  [self.lock lock];
  NSTouchBar* bar = [[NSTouchBar alloc] init];
  [self setupTouchBar:bar];
  [bar setDelegate:self];
  [bar autorelease];
  [self.lock unlock];
  return bar;
}

- (nullable NSTouchBarItem *)touchBar:(NSTouchBar *)touchBar makeItemForIdentifier:(NSTouchBarItemIdentifier)identifier {
  if ([identifier isEqual:@""]) {
    return nil;
  }

  NSTouchBarItem* item = nil;
  [self.lock lock];
  NSDictionary* data = [[self.goData objectForKey:@"Items"] objectForKey:identifier];

  // TODO: finish
  if ([identifier hasPrefix:prefixButton]) {
    NSButtonTouchBarItem* myItem = [[[NSButtonTouchBarItem alloc] initWithIdentifier:identifier] autorelease];
    [self updateWidgetButton:myItem withData:data];
    item = myItem;

  } else if ([identifier hasPrefix:prefixLabel]) {
    NSCustomTouchBarItem* myItem = [[[NSCustomTouchBarItem alloc] initWithIdentifier:identifier] autorelease];
    [self updateWidgetLabel:myItem withData:data];
    item = myItem;

  } else if ([identifier hasPrefix:prefixPopover]) {
    NSPopoverTouchBarItem* myItem = [[[NSPopoverTouchBarItem alloc] initWithIdentifier:identifier] autorelease];
    [self updateWidgetPopover:myItem touchBar:[[NSTouchBar alloc] init] withData:data];
    item = myItem;

  } else if ([identifier hasPrefix:prefixSlider]) {
    NSSliderTouchBarItem* myItem = [[[NSSliderTouchBarItem alloc] initWithIdentifier:identifier] autorelease];
    [self updateWidgetSlider:myItem withData:data];
    myItem.slider.doubleValue = [[data valueForKeyPath:@"StartValue"] doubleValue];
    item = myItem;

  } else {
    NSLog(@"warning: unsupported identifier %@ with %@", identifier, data);
  }

  [self.lock unlock];
  return item;
}

- (void)updateItem:(NSTouchBar *)touchBar withIdentifier:(NSTouchBarItemIdentifier)identifier {
  if ([identifier isEqual:@""] || ![identifier hasPrefix:prefixWidget]) {
    return;
  }
  id item = [touchBar itemForIdentifier:identifier];
  if (item == nil) {
    return;
  }

  [self.lock lock];
  NSDictionary* data = [[self.goData objectForKey:@"Items"] objectForKey:identifier];
  dispatch_async(dispatch_get_main_queue(), ^{
    // TODO: finish
    if ([identifier hasPrefix:prefixButton]) {
      [self updateWidgetButton:item withData:data];
    } else if ([identifier hasPrefix:prefixLabel]) {
      [self updateWidgetLabel:item withData:data];
    } else if ([identifier hasPrefix:prefixPopover]) {
      NSPopoverTouchBarItem* pop = (NSPopoverTouchBarItem*) item;
      NSTouchBar* sub = pop.popoverTouchBar != nil ? pop.popoverTouchBar : pop.pressAndHoldTouchBar;
      [self updateWidgetPopover:item touchBar:sub withData:data];
      for (id child in [data objectForKey:@"Bar"]) {
        [self updateItem:sub withIdentifier:[self transformIdentifier:child]];
      }
    } else if ([identifier hasPrefix:prefixSlider]) {
      [self updateWidgetSlider:item withData:data];
    } else {
      NSLog(@"warning: unknown identifier %@ with %@", identifier, data);
    }
  });
  [self.lock unlock];
}

- (void)updateItems:(NSTouchBar *)touchBar {
  [self updateItem:touchBar withIdentifier:touchBar.escapeKeyReplacementItemIdentifier];
  [self updateItem:touchBar withIdentifier:touchBar.principalItemIdentifier];
  for (id item in touchBar.itemIdentifiers) {
    [self updateItem:touchBar withIdentifier:item];
  }
}

- (void)updateWidgetCore:(NSTouchBarItem*)item withData:(NSDictionary*)data {
  item.visibilityPriority = [[data objectForKey:@"Priority"] floatValue];
}

- (void)updateWidgetButton:(NSButtonTouchBarItem*)item withData:(NSDictionary*)data {
  [self updateWidgetCore:item withData:data];

  item.title = [data valueForKeyPath:@"Title"];
  item.image = [self transformImage:[data valueForKeyPath:@"Image"]];
  item.target = self;
  item.action = @selector(buttonAction:);
  item.enabled = [[data valueForKeyPath:@"Disabled"] intValue] == 0;
  item.bezelColor = [self transformColor:[data valueForKeyPath:@"BezelColor"]];
}

- (void)buttonAction:(id)sender {
  NSString* identifier = ((NSButtonTouchBarItem*) sender).identifier;
  const char * event = [[NSString stringWithFormat:@"{\"Kind\":\"button\",\"Target\":\"%@\"}", identifier] UTF8String];
  self.handler((char*) event);
}

- (void)updateWidgetPopover:(NSPopoverTouchBarItem*)item touchBar:(NSTouchBar*)sub withData:(NSDictionary*)data {
  [self updateWidgetCore:item withData:data];

  item.collapsedRepresentationLabel = [data valueForKeyPath:@"CollapsedText"];
  item.collapsedRepresentationImage = [self transformImage:[data valueForKeyPath:@"CollapsedImage"]];
  item.popoverTouchBar = sub;
  if ([[data valueForKeyPath:@"PressAndHold"] intValue] == 1) {
    item.pressAndHoldTouchBar = sub;
  } else {
    item.pressAndHoldTouchBar = nil;
  }
  sub.defaultItemIdentifiers = [self mapIdentifiers:[data objectForKey:@"Bar"]];
  sub.principalItemIdentifier = [self transformIdentifier:[data objectForKey:@"Principal"]];
  sub.delegate = self;
}

- (void)updateWidgetSlider:(NSSliderTouchBarItem*)item withData:(NSDictionary*)data {
  [self updateWidgetCore:item withData:data];

  // FIXME: weird NSLayoutConstraintNumberExceedsLimit warning
  item.label = [data valueForKeyPath:@"Label"];
  NSImage* minimumAccessory = [self transformImage:[data valueForKeyPath:@"MinimumAccessory"]];
  if (minimumAccessory != nil) {
    item.minimumValueAccessory = [NSSliderAccessory accessoryWithImage:minimumAccessory];
  } else {
    item.minimumValueAccessory = nil;
  }
  NSImage* maximumAccessory = [self transformImage:[data valueForKeyPath:@"MaximumAccessory"]];
  if (maximumAccessory != nil) {
    item.maximumValueAccessory = [NSSliderAccessory accessoryWithImage:maximumAccessory];
  } else {
    item.maximumValueAccessory = nil;
  }
  NSString* accessoryWidth = [data valueForKeyPath:@"AccessoryWidth"];
  if ([accessoryWidth isEqual:@"wide"]) {
    item.valueAccessoryWidth = NSSliderAccessoryWidthWide;
  } else {
    item.valueAccessoryWidth = NSSliderAccessoryWidthDefault;
  }
  item.slider.minValue = [[data valueForKeyPath:@"MinimumValue"] doubleValue];
  item.slider.maxValue = [[data valueForKeyPath:@"MaximumValue"] doubleValue];
  item.target = self;
  item.action = @selector(sliderAction:);
}

- (void)sliderAction:(NSSliderTouchBarItem*) sender {
  NSString* identifier = sender.identifier;
  const char * event = [[NSString
    stringWithFormat:@"{\"Kind\":\"slider\",\"Target\":\"%@\",\"Data\":%f}",
    identifier,
    sender.slider.doubleValue
  ] UTF8String];
  self.handler((char*) event);
}

- (void)updateWidgetLabel:(NSCustomTouchBarItem*)item withData:(NSDictionary*)data {
  [self updateWidgetCore:item withData:data];

  NSString* text = [data valueForKeyPath:@"Content.Text"];
  NSString* image = [data valueForKeyPath:@"Content.Image"];
  if (text != nil) {
    NSTextField* view = [NSTextField labelWithString:text];
    view.textColor = [self transformColor:[data valueForKeyPath:@"Content.Color"]];;
    [item setView:view];
  } else if (image != nil) {
    NSImageView* view = [NSImageView imageViewWithImage: [self transformImage:image]];
    [item setView:view];
  } else {
    NSLog(@"warning: label with invalid data %@", data);
  }
}

- (NSTouchBarItemIdentifier) transformIdentifier:(NSString*) name {
  NSTouchBarItemIdentifier standard = [self.identifierMapping objectForKey:name];
  if (standard != nil) {
    return standard;
  }
  return name;
}

- (NSImage*) transformImage:(NSString*)name {
  if (name == nil || name == (id)[NSNull null]) {
    return nil;
  }
  NSImageName standard = [self.imageMapping objectForKey:name];
  if (standard != nil) {
    return [NSImage imageNamed:standard];
  }
  NSImage* sf = [NSImage imageWithSystemSymbolName:name accessibilityDescription:name];
  if (sf == nil) {
    NSLog(@"warning: could not find SF Symbols for %@", name);
  }
  return sf;
}

- (NSColor*) transformColor:(NSDictionary*)details {
  if (details == nil || details == (id)[NSNull null]) {
    return nil;
  }
  if ([details isKindOfClass:[NSString class]]) {
    NSString* strDetails = (id)details;
    int length = [strDetails length];

    if ([self.hexRegex firstMatchInString:strDetails options:0 range:NSMakeRange(0, length)] == nil) {
      // no idea
      NSLog(@"unsupported color: %@", strDetails);
      return nil;
    }

    // hex string
    if ([strDetails hasPrefix:@"#"]) {
      strDetails = [strDetails substringWithRange:NSMakeRange(1, length - 1)];
    }
    NSScanner* scanner = [NSScanner scannerWithString:strDetails];
    unsigned int hexColor;
    if (![scanner scanHexInt:&hexColor]) {
      NSLog(@"invalid hex color: %@", strDetails);
      return nil;
    }
    return [NSColor
      colorWithSRGBRed:(CGFloat)((hexColor >> 16) & 0xff) / 255.0
      green:(CGFloat)((hexColor >> 8) & 0xff) / 255.0
      blue:(CGFloat)((hexColor >> 0) & 0xff) / 255.0
      alpha:1.0
    ];
  }
  // rgba object
  return [NSColor
    colorWithSRGBRed:[[details objectForKey:@"Red"] doubleValue]
    green:[[details objectForKey:@"Green"] doubleValue]
    blue:[[details objectForKey:@"Blue"] doubleValue]
    alpha:[[details objectForKey:@"Alpha"] doubleValue]
  ];
}

- (NSError*) initConst {
  NSError* error = nil;
  self.hexRegex = [NSRegularExpression
    regularExpressionWithPattern:@"^#?[a-f0-9]{6}$"
    options:NSRegularExpressionCaseInsensitive
    error:&error
  ];
  if (self.hexRegex == nil) {
    if (error != nil) {
      return error;
    }
    return [[[NSError alloc] initWithDomain:@"com.energy.touchbar.go" code:1 userInfo:@{@"Error reason": @"cannot make regex"}] autorelease];
  }
  return nil;
}

- (void) initMapping {
  self.identifierMapping = @{
    standardOtherItems: NSTouchBarItemIdentifierOtherItemsProxy,
    standardSpaceSmall: NSTouchBarItemIdentifierFixedSpaceSmall,
    standardSpaceLarge: NSTouchBarItemIdentifierFixedSpaceLarge,
    standardSpaceFlexible: NSTouchBarItemIdentifierFlexibleSpace,
    standardCandidateList: NSTouchBarItemIdentifierCandidateList,
    standardCharacterPicker: NSTouchBarItemIdentifierCharacterPicker,
    standardTextFormat: NSTouchBarItemIdentifierTextFormat,
    standardTextAlignment: NSTouchBarItemIdentifierTextAlignment,
    standardTextColorPicker: NSTouchBarItemIdentifierTextColorPicker,
    standardTextList: NSTouchBarItemIdentifierTextList,
    standardTextStyle: NSTouchBarItemIdentifierTextStyle,
  };
  self.imageMapping = @{
    @"TBAddDetailTemplate": NSImageNameTouchBarAddDetailTemplate,
    @"TBAddTemplate": NSImageNameTouchBarAddTemplate,
    @"TBAlarmTemplate": NSImageNameTouchBarAlarmTemplate,
    @"TBAudioInputMuteTemplate": NSImageNameTouchBarAudioInputMuteTemplate,
    @"TBAudioInputTemplate": NSImageNameTouchBarAudioInputTemplate,
    @"TBAudioOutputMuteTemplate": NSImageNameTouchBarAudioOutputMuteTemplate,
    @"TBAudioOutputVolumeHighTemplate": NSImageNameTouchBarAudioOutputVolumeHighTemplate,
    @"TBAudioOutputVolumeLowTemplate": NSImageNameTouchBarAudioOutputVolumeLowTemplate,
    @"TBAudioOutputVolumeMediumTemplate": NSImageNameTouchBarAudioOutputVolumeMediumTemplate,
    @"TBAudioOutputVolumeOffTemplate": NSImageNameTouchBarAudioOutputVolumeOffTemplate,
    @"TBBookmarksTemplate": NSImageNameTouchBarBookmarksTemplate,
    @"TBColorPickerFill": NSImageNameTouchBarColorPickerFill,
    @"TBColorPickerFont": NSImageNameTouchBarColorPickerFont,
    @"TBColorPickerStroke": NSImageNameTouchBarColorPickerStroke,
    @"TBCommunicationAudioTemplate": NSImageNameTouchBarCommunicationAudioTemplate,
    @"TBCommunicationVideoTemplate": NSImageNameTouchBarCommunicationVideoTemplate,
    @"TBComposeTemplate": NSImageNameTouchBarComposeTemplate,
    @"TBDeleteTemplate": NSImageNameTouchBarDeleteTemplate,
    @"TBDownloadTemplate": NSImageNameTouchBarDownloadTemplate,
    @"TBEnterFullScreenTemplate": NSImageNameTouchBarEnterFullScreenTemplate,
    @"TBExitFullScreenTemplate": NSImageNameTouchBarExitFullScreenTemplate,
    @"TBFastForwardTemplate": NSImageNameTouchBarFastForwardTemplate,
    @"TBFolderCopyToTemplate": NSImageNameTouchBarFolderCopyToTemplate,
    @"TBFolderMoveToTemplate": NSImageNameTouchBarFolderMoveToTemplate,
    @"TBFolderTemplate": NSImageNameTouchBarFolderTemplate,
    @"TBGetInfoTemplate": NSImageNameTouchBarGetInfoTemplate,
    @"TBGoBackTemplate": NSImageNameTouchBarGoBackTemplate,
    @"TBGoDownTemplate": NSImageNameTouchBarGoDownTemplate,
    @"TBGoForwardTemplate": NSImageNameTouchBarGoForwardTemplate,
    @"TBGoUpTemplate": NSImageNameTouchBarGoUpTemplate,
    @"TBHistoryTemplate": NSImageNameTouchBarHistoryTemplate,
    @"TBIconViewTemplate": NSImageNameTouchBarIconViewTemplate,
    @"TBListViewTemplate": NSImageNameTouchBarListViewTemplate,
    @"TBMailTemplate": NSImageNameTouchBarMailTemplate,
    @"TBNewFolderTemplate": NSImageNameTouchBarNewFolderTemplate,
    @"TBNewMessageTemplate": NSImageNameTouchBarNewMessageTemplate,
    @"TBOpenInBrowserTemplate": NSImageNameTouchBarOpenInBrowserTemplate,
    @"TBPauseTemplate": NSImageNameTouchBarPauseTemplate,
    @"TBPlayheadTemplate": NSImageNameTouchBarPlayheadTemplate,
    @"TBPlayPauseTemplate": NSImageNameTouchBarPlayPauseTemplate,
    @"TBPlayTemplate": NSImageNameTouchBarPlayTemplate,
    @"TBQuickLookTemplate": NSImageNameTouchBarQuickLookTemplate,
    @"TBRecordStartTemplate": NSImageNameTouchBarRecordStartTemplate,
    @"TBRecordStopTemplate": NSImageNameTouchBarRecordStopTemplate,
    @"TBRefreshTemplate": NSImageNameTouchBarRefreshTemplate,
    @"TBRewindTemplate": NSImageNameTouchBarRewindTemplate,
    @"TBRotateLeftTemplate": NSImageNameTouchBarRotateLeftTemplate,
    @"TBRotateRightTemplate": NSImageNameTouchBarRotateRightTemplate,
    @"TBSearchTemplate": NSImageNameTouchBarSearchTemplate,
    @"TBShareTemplate": NSImageNameTouchBarShareTemplate,
    @"TBSidebarTemplate": NSImageNameTouchBarSidebarTemplate,
    @"TBSkipAhead15SecondsTemplate": NSImageNameTouchBarSkipAhead15SecondsTemplate,
    @"TBSkipAhead30SecondsTemplate": NSImageNameTouchBarSkipAhead30SecondsTemplate,
    @"TBSkipAheadTemplate": NSImageNameTouchBarSkipAheadTemplate,
    @"TBSkipBack15SecondsTemplate": NSImageNameTouchBarSkipBack15SecondsTemplate,
    @"TBSkipBack30SecondsTemplate": NSImageNameTouchBarSkipBack30SecondsTemplate,
    @"TBSkipBackTemplate": NSImageNameTouchBarSkipBackTemplate,
    @"TBSkipToEndTemplate": NSImageNameTouchBarSkipToEndTemplate,
    @"TBSkipToStartTemplate": NSImageNameTouchBarSkipToStartTemplate,
    @"TBSlideshowTemplate": NSImageNameTouchBarSlideshowTemplate,
    @"TBTagIconTemplate": NSImageNameTouchBarTagIconTemplate,
    @"TBTextBoldTemplate": NSImageNameTouchBarTextBoldTemplate,
    @"TBTextBoxTemplate": NSImageNameTouchBarTextBoxTemplate,
    @"TBTextCenterAlignTemplate": NSImageNameTouchBarTextCenterAlignTemplate,
    @"TBTextItalicTemplate": NSImageNameTouchBarTextItalicTemplate,
    @"TBTextJustifiedAlignTemplate": NSImageNameTouchBarTextJustifiedAlignTemplate,
    @"TBTextLeftAlignTemplate": NSImageNameTouchBarTextLeftAlignTemplate,
    @"TBTextListTemplate": NSImageNameTouchBarTextListTemplate,
    @"TBTextRightAlignTemplate": NSImageNameTouchBarTextRightAlignTemplate,
    @"TBTextStrikethroughTemplate": NSImageNameTouchBarTextStrikethroughTemplate,
    @"TBTextUnderlineTemplate": NSImageNameTouchBarTextUnderlineTemplate,
    @"TBUserAddTemplate": NSImageNameTouchBarUserAddTemplate,
    @"TBUserGroupTemplate": NSImageNameTouchBarUserGroupTemplate,
    @"TBUserTemplate": NSImageNameTouchBarUserTemplate,
  };
}
@end
