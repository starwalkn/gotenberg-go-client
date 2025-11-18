package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

type ChromiumService struct {
	client *Client
}

type LibreOfficeService struct {
	client *Client
}

type PDFEnginesService struct {
	client *Client
}

func (s *ChromiumService) HTML(index document.Document) *HTMLRequest {
	return &HTMLRequest{
		index:           index,
		assets:          []document.Document{},
		embeds:          []document.Document{},
		chromiumRequest: newChromiumRequest(),
		client:          s.client,
	}
}

func (s *ChromiumService) Markdown(index document.Document, markdowns ...document.Document) *MarkdownRequest {
	return &MarkdownRequest{
		index:           index,
		markdowns:       markdowns,
		assets:          []document.Document{},
		embeds:          []document.Document{},
		chromiumRequest: newChromiumRequest(),
		client:          s.client,
	}
}

func (s *ChromiumService) URL(url string) *URLRequest {
	crr := newChromiumRequest()
	crr.fields[fieldURL] = url

	return &URLRequest{
		embeds:          []document.Document{},
		chromiumRequest: crr,
		client:          s.client,
	}
}

func (s *LibreOfficeService) HTML(docs ...document.Document) *LibreOfficeRequest {
	return &LibreOfficeRequest{
		docs:        docs,
		embeds:      []document.Document{},
		baseRequest: newBaseRequest(),
		client:      s.client,
	}
}

func (s *PDFEnginesService) Merge(pdfs ...document.Document) *MergeRequest {
	return &MergeRequest{
		pdfs:        pdfs,
		embeds:      []document.Document{},
		baseRequest: newBaseRequest(),
		client:      s.client,
	}
}

func (s *PDFEnginesService) Flatten(pdfs ...document.Document) *FlattenRequest {
	return &FlattenRequest{
		pdfs:        pdfs,
		baseRequest: newBaseRequest(),
		client:      s.client,
	}
}

func (s *PDFEnginesService) WriteMetadata(pdfs ...document.Document) *WriteMetadataRequest {
	return &WriteMetadataRequest{
		pdfs:        pdfs,
		embeds:      []document.Document{},
		baseRequest: newBaseRequest(),
		client:      s.client,
	}
}

func (s *PDFEnginesService) ReadMetadata(pdfs ...document.Document) *ReadMetadataRequest {
	return &ReadMetadataRequest{
		pdfs:        pdfs,
		embeds:      []document.Document{},
		baseRequest: newBaseRequest(),
		client:      s.client,
	}
}

func (s *PDFEnginesService) Encrypt(userPassword, ownerPassword string, pdfs ...document.Document) *EncryptRequest {
	br := newBaseRequest()
	br.fields[fieldUserPassword] = userPassword

	if ownerPassword != "" {
		br.fields[fieldOwnerPassword] = ownerPassword
	}

	return &EncryptRequest{
		pdfs:        pdfs,
		baseRequest: br,
		client:      s.client,
	}
}

func (s *PDFEnginesService) Embed(pdfs, embeds []document.Document) *EmbedRequest {
	return &EmbedRequest{
		pdfs:        pdfs,
		embeds:      embeds,
		baseRequest: newBaseRequest(),
		client:      s.client,
	}
}

func (s *PDFEnginesService) SplitIntervals(pdfs ...document.Document) *SplitIntervalsRequest {
	br := newBaseRequest()
	br.fields[fieldSplitMode] = splitModeIntervals

	return &SplitIntervalsRequest{
		pdfs:        pdfs,
		embeds:      []document.Document{},
		baseRequest: br,
		client:      s.client,
	}
}

func (s *PDFEnginesService) SplitPages(pdfs ...document.Document) *SplitPagesRequest {
	br := newBaseRequest()
	br.fields[fieldSplitMode] = splitModePages

	return &SplitPagesRequest{
		pdfs:        pdfs,
		embeds:      []document.Document{},
		baseRequest: br,
		client:      s.client,
	}
}
