package utils

import (
	"context"
	"fmt"

	ImageKit "github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

func ImageHandler(file interface{}, filename string, folder string) (string, error) {
	var ctx = context.Background()
	// Using keys in argument
	ik := ImageKit.NewFromParams(ImageKit.NewParams{
		PrivateKey:  "private_ZQ3wyHkY98InMeoVgUCc7kORoAo=",
		PublicKey:   "public_+LsWTN2IGXkaGmgXD8PpE/n7HFo=",
		UrlEndpoint: "https://ik.imagekit.io/blogemyu",
	})

	resp, _ := ik.Uploader.Upload(ctx, file, uploader.UploadParam{
		Folder:   fmt.Sprintf("initoko/%s", folder),
		FileName: filename,
	})

	return resp.Data.Url, nil

}
