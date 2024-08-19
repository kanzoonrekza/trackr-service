package types

import "mime/multipart"

type FormDataFields map[string]string
type FormDataFile map[string][]*multipart.FileHeader
