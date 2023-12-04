package barbuilder

// Image represents an image usable by macOS
type Image interface {
	isAnImage()
}

// SFSymbol represents a standard Apple symbol compatible with the San Francisco font
// See https://developer.apple.com/sf-symbols/ for full list
type SFSymbol string

var _ Image = SFSymbol("")

func (me SFSymbol) isAnImage() {}

// TBSymbol represents a standard TouchBar icon
// Note: if possible you should use `SFSymbol` instead
// See https://developer.apple.com/design/human-interface-guidelines/inputs/touch-bar/#interface-icons for full list and mapping to `SFSymbol`
type TBSymbol string

var _ Image = TBSymbol("")

func (me TBSymbol) isAnImage() {}

const (
	TBAddDetailTemplate               TBSymbol = "TBAddDetailTemplate"
	TBAddTemplate                     TBSymbol = "TBAddTemplate"
	TBAlarmTemplate                   TBSymbol = "TBAlarmTemplate"
	TBAudioInputMuteTemplate          TBSymbol = "TBAudioInputMuteTemplate"
	TBAudioInputTemplate              TBSymbol = "TBAudioInputTemplate"
	TBAudioOutputMuteTemplate         TBSymbol = "TBAudioOutputMuteTemplate"
	TBAudioOutputVolumeHighTemplate   TBSymbol = "TBAudioOutputVolumeHighTemplate"
	TBAudioOutputVolumeLowTemplate    TBSymbol = "TBAudioOutputVolumeLowTemplate"
	TBAudioOutputVolumeMediumTemplate TBSymbol = "TBAudioOutputVolumeMediumTemplate"
	TBAudioOutputVolumeOffTemplate    TBSymbol = "TBAudioOutputVolumeOffTemplate"
	TBBookmarksTemplate               TBSymbol = "TBBookmarksTemplate"
	TBColorPickerFill                 TBSymbol = "TBColorPickerFill"
	TBColorPickerFont                 TBSymbol = "TBColorPickerFont"
	TBColorPickerStroke               TBSymbol = "TBColorPickerStroke"
	TBCommunicationAudioTemplate      TBSymbol = "TBCommunicationAudioTemplate"
	TBCommunicationVideoTemplate      TBSymbol = "TBCommunicationVideoTemplate"
	TBComposeTemplate                 TBSymbol = "TBComposeTemplate"
	TBDeleteTemplate                  TBSymbol = "TBDeleteTemplate"
	TBDownloadTemplate                TBSymbol = "TBDownloadTemplate"
	TBEnterFullScreenTemplate         TBSymbol = "TBEnterFullScreenTemplate"
	TBExitFullScreenTemplate          TBSymbol = "TBExitFullScreenTemplate"
	TBFastForwardTemplate             TBSymbol = "TBFastForwardTemplate"
	TBFolderCopyToTemplate            TBSymbol = "TBFolderCopyToTemplate"
	TBFolderMoveToTemplate            TBSymbol = "TBFolderMoveToTemplate"
	TBFolderTemplate                  TBSymbol = "TBFolderTemplate"
	TBGetInfoTemplate                 TBSymbol = "TBGetInfoTemplate"
	TBGoBackTemplate                  TBSymbol = "TBGoBackTemplate"
	TBGoDownTemplate                  TBSymbol = "TBGoDownTemplate"
	TBGoForwardTemplate               TBSymbol = "TBGoForwardTemplate"
	TBGoUpTemplate                    TBSymbol = "TBGoUpTemplate"
	TBHistoryTemplate                 TBSymbol = "TBHistoryTemplate"
	TBIconViewTemplate                TBSymbol = "TBIconViewTemplate"
	TBListViewTemplate                TBSymbol = "TBListViewTemplate"
	TBMailTemplate                    TBSymbol = "TBMailTemplate"
	TBNewFolderTemplate               TBSymbol = "TBNewFolderTemplate"
	TBNewMessageTemplate              TBSymbol = "TBNewMessageTemplate"
	TBOpenInBrowserTemplate           TBSymbol = "TBOpenInBrowserTemplate"
	TBPauseTemplate                   TBSymbol = "TBPauseTemplate"
	TBPlayheadTemplate                TBSymbol = "TBPlayheadTemplate"
	TBPlayPauseTemplate               TBSymbol = "TBPlayPauseTemplate"
	TBPlayTemplate                    TBSymbol = "TBPlayTemplate"
	TBQuickLookTemplate               TBSymbol = "TBQuickLookTemplate"
	TBRecordStartTemplate             TBSymbol = "TBRecordStartTemplate"
	TBRecordStopTemplate              TBSymbol = "TBRecordStopTemplate"
	TBRefreshTemplate                 TBSymbol = "TBRefreshTemplate"
	TBRewindTemplate                  TBSymbol = "TBRewindTemplate"
	TBRotateLeftTemplate              TBSymbol = "TBRotateLeftTemplate"
	TBRotateRightTemplate             TBSymbol = "TBRotateRightTemplate"
	TBSearchTemplate                  TBSymbol = "TBSearchTemplate"
	TBShareTemplate                   TBSymbol = "TBShareTemplate"
	TBSidebarTemplate                 TBSymbol = "TBSidebarTemplate"
	TBSkipAhead15SecondsTemplate      TBSymbol = "TBSkipAhead15SecondsTemplate"
	TBSkipAhead30SecondsTemplate      TBSymbol = "TBSkipAhead30SecondsTemplate"
	TBSkipAheadTemplate               TBSymbol = "TBSkipAheadTemplate"
	TBSkipBack15SecondsTemplate       TBSymbol = "TBSkipBack15SecondsTemplate"
	TBSkipBack30SecondsTemplate       TBSymbol = "TBSkipBack30SecondsTemplate"
	TBSkipBackTemplate                TBSymbol = "TBSkipBackTemplate"
	TBSkipToEndTemplate               TBSymbol = "TBSkipToEndTemplate"
	TBSkipToStartTemplate             TBSymbol = "TBSkipToStartTemplate"
	TBSlideshowTemplate               TBSymbol = "TBSlideshowTemplate"
	TBTagIconTemplate                 TBSymbol = "TBTagIconTemplate"
	TBTextBoldTemplate                TBSymbol = "TBTextBoldTemplate"
	TBTextBoxTemplate                 TBSymbol = "TBTextBoxTemplate"
	TBTextCenterAlignTemplate         TBSymbol = "TBTextCenterAlignTemplate"
	TBTextItalicTemplate              TBSymbol = "TBTextItalicTemplate"
	TBTextJustifiedAlignTemplate      TBSymbol = "TBTextJustifiedAlignTemplate"
	TBTextLeftAlignTemplate           TBSymbol = "TBTextLeftAlignTemplate"
	TBTextListTemplate                TBSymbol = "TBTextListTemplate"
	TBTextRightAlignTemplate          TBSymbol = "TBTextRightAlignTemplate"
	TBTextStrikethroughTemplate       TBSymbol = "TBTextStrikethroughTemplate"
	TBTextUnderlineTemplate           TBSymbol = "TBTextUnderlineTemplate"
	TBUserAddTemplate                 TBSymbol = "TBUserAddTemplate"
	TBUserGroupTemplate               TBSymbol = "TBUserGroupTemplate"
	TBUserTemplate                    TBSymbol = "TBUserTemplate"
)
