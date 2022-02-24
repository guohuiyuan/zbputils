package pool

import (
	"errors"

	"github.com/FloatTech/zbputils/ctxext"
	"github.com/FloatTech/zbputils/file"
	"github.com/FloatTech/zbputils/web"
	"github.com/sirupsen/logrus"
	"github.com/wdvxdr1123/ZeroBot/message"
)

// SendImageFromPool ...
func SendImageFromPool(imgname, imgpath string, genimg func() error, send ctxext.NoCtxSendMsg, get ctxext.NoCtxGetMsg) error {
	m, err := GetImage(imgname)
	if err != nil {
		logrus.Debugln("[ctxext.img]", err)
		if file.IsNotExist(imgpath) {
			err = genimg()
			if err != nil {
				return err
			}
		}
		m.SetFile(file.BOTPATH + "/" + imgpath)
		hassent, err := m.Push(send, get)
		if hassent {
			return nil
		}
		if err != nil {
			return err
		}
	}
	// 发送图片
	img := message.Image(m.String())
	id := send(message.Message{img})
	if id == 0 {
		id = send(message.Message{img.Add("cache", "0")})
		if id == 0 {
			data, err := web.GetData(m.String())
			if err != nil {
				return err
			}
			id = send(message.Message{message.ImageBytes(data)})
			if id == 0 {
				return errors.New("图片发送失败，可能被风控了~")
			}
		}
	}
	return nil
}