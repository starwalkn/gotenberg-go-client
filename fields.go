package gotenberg

type formField string

// Common property.
const (
	fieldMetadata     formField = "metadata"
	fieldDownloadFrom formField = "downloadFrom"
)

// URL request property.
const (
	fieldURL formField = "url"
)

// Chromium request configuration.
const (
	fieldChromiumWaitDelay                     formField = "waitDelay"
	fieldChromiumWaitForExpression             formField = "waitForExpression"
	fieldChromiumEmulatedMediaType             formField = "emulatedMediaType"
	fieldChromiumCookies                       formField = "cookies"
	fieldChromiumUserAgent                     formField = "userAgent"
	fieldChromiumExtraHTTPHeaders              formField = "extraHttpHeaders"
	fieldChromiumFailOnHTTPStatusCodes         formField = "failOnHttpStatusCodes"
	fieldChromiumFailOnResourceHTTPStatusCodes formField = "failOnResourceHttpStatusCodes"
	fieldChromiumFailOnConsoleExceptions       formField = "failOnConsoleExceptions"
	fieldChromiumFailOnResourceLoadingFailed   formField = "failOnResourceLoadingFailed"
	fieldChromiumSkipNetworkIdleEvent          formField = "skipNetworkIdleEvent"
	fieldChromiumGenerateTaggedPDF             formField = "generateTaggedPdf"
)

// PDF page property.
const (
	fieldChromiumPaperWidth              formField = "paperWidth"
	fieldChromiumPaperHeight             formField = "paperHeight"
	fieldChromiumMarginTop               formField = "marginTop"
	fieldChromiumMarginBottom            formField = "marginBottom"
	fieldChromiumMarginLeft              formField = "marginLeft"
	fieldChromiumMarginRight             formField = "marginRight"
	fieldChromiumLandscapeChrome         formField = "landscape"
	fieldChromiumNativePageRanges        formField = "nativePageRanges"
	fieldChromiumScale                   formField = "scale"
	fieldChromiumSinglePage              formField = "singlePage"
	fieldChromiumPreferCSSPageSize       formField = "preferCssPageSize"
	fieldChromiumPrintBackground         formField = "printBackground"
	fieldChromiumOmitBackground          formField = "omitBackground"
	fieldChromiumGenerateDocumentOutline formField = "generateDocumentOutline"
)

// Screenshot property.
const (
	fieldScreenshotWidth            formField = "width"
	fieldScreenshotHeight           formField = "height"
	fieldScreenshotClip             formField = "clip"
	fieldScreenshotFormat           formField = "format"
	fieldScreenshotQuality          formField = "quality"
	fieldScreenshotOptimizeForSpeed formField = "optimizeForSpeed"
)

// Office request property.
const (
	fieldOfficePassword                        formField = "password"
	fieldOfficeLandscape                       formField = "landscape"
	fieldOfficeNativePageRanges                formField = "nativePageRanges"
	fieldOfficeExportFormFields                formField = "exportFormFields"
	fieldOfficeAllowDuplicateFieldNames        formField = "allowDuplicateFieldNames"
	fieldOfficeExportBookmarks                 formField = "exportBookmarks"
	fieldOfficeExportBookmarksToPdfDestination formField = "exportBookmarksToPdfDestination"
	fieldOfficeExportPlaceholders              formField = "exportPlaceholders"
	fieldOfficeExportNotes                     formField = "exportNotes"
	fieldOfficeExportNotesPages                formField = "exportNotesPages"
	fieldOfficeExportOnlyNotesPages            formField = "exportOnlyNotesPages"
	fieldOfficeExportNotesInMargin             formField = "exportNotesInMargin"
	fieldOfficeConvertOooTargetToPdfTarget     formField = "convertOooTargetToPdfTarget"
	fieldOfficeExportLinksRelativeFsys         formField = "exportLinksRelativeFsys"
	fieldOfficeExportHiddenSlides              formField = "exportHiddenSlides"
	fieldOfficeSkipEmptyPages                  formField = "skipEmptyPages"
	fieldOfficeAddOriginalDocumentAsStream     formField = "addOriginalDocumentAsStream"
	fieldOfficeSinglePageSheets                formField = "singlePageSheets"
	fieldOfficeLosslessImageCompression        formField = "losslessImageCompression"
	fieldOfficeQuality                         formField = "quality"
	fieldOfficeReduceImageResolution           formField = "reduceImageResolution"
	fieldOfficeMaxImageResolution              formField = "maxImageResolution"
	fieldOfficeMerge                           formField = "merge"
	fieldOfficePdfA                            formField = "pdfa"
	fieldOfficePdfUa                           formField = "pdfua"
	fieldOfficeFlatten                         formField = "flatten"
	fieldOfficeUpdateIndexes                   formField = "updateIndexes"
)

// Merge request property.
const (
	fieldMergePdfA    formField = "pdfa"
	fieldMergePdfUA   formField = "pdfua"
	fieldMergeFlatten formField = "flatten"
)

const (
	fieldSplitMode    = "splitMode"
	fieldSplitSpan    = "splitSpan"
	fieldSplitUnify   = "splitUnify"
	fieldSplitFlatten = "flatten"
)

// Split mode.
const (
	splitModePages     = "pages"
	splitModeIntervals = "intervals"
)
