package types

import "mime/multipart"

type FormDataFile map[string][]*multipart.FileHeader
