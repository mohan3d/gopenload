package openload

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func c() *Client {
	return New("LOGIN", "KEY", nil)
}

func TestAccountInfo(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/account/info").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"extid":"extuserid","email":"jeff@openload.io","signup_at":"2015-01-09 23:59:54","storage_left":-1,"storage_used":"32922117680","traffic":{"left":-1,"used_24h":0},"balance":0}}`)

	info, err := c().AccountInfo()

	assert.Nil(t, err)
	assert.EqualValues(t, "extuserid", info.Extid)
	assert.EqualValues(t, "jeff@openload.io", info.Email)
	assert.EqualValues(t, "2015-01-09 23:59:54", info.SignupAt)
	assert.EqualValues(t, -1, info.StorageLeft)
	assert.EqualValues(t, "32922117680", info.StorageUsed)
	assert.EqualValues(t, -1, info.Traffic.Left)
	assert.EqualValues(t, 0, info.Traffic.Used24H)
	assert.EqualValues(t, 0, info.Balance)
}

func TestDownloadTicket(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/dlticket").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"ticket":"72fA-_Lq8Ak~~1440353112~n~~0~nXtN3RI-nsEa28Iq","captcha_url":"https://openload.co/dlcaptcha/b92eY_nfjV4.png","captcha_w":140,"captcha_h":70,"wait_time":10,"valid_until":"2015-08-23 18:20:13"}}`)

	ticket, err := c().DownloadTicket("<FILE_ID>")

	assert.Nil(t, err)
	assert.EqualValues(t, "72fA-_Lq8Ak~~1440353112~n~~0~nXtN3RI-nsEa28Iq", ticket.Ticket)
	assert.EqualValues(t, "https://openload.co/dlcaptcha/b92eY_nfjV4.png", ticket.CaptchaURL)
	assert.EqualValues(t, 140, ticket.CaptchaW)
	assert.EqualValues(t, 70, ticket.CaptchaH)
	assert.EqualValues(t, 10, ticket.WaitTime)
	assert.EqualValues(t, "2015-08-23 18:20:13", ticket.ValidUntil)
}

func TestDownloadLink(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/dl").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"name":"The quick brown fox.txt","size":12345,"sha1":"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12","content_type":"plain/text","upload_at":"2011-01-26 13:33:37","url":"https://abvzps.example.com/dl/l/4spxX_-cSO4/The+quick+brown+fox.txt","token":"4spxX_-cSO4"}}`)

	link, err := c().DownloadLink("<FILE_ID>", "4spxX_-cSO4", "captchas suck")

	assert.Nil(t, err)
	assert.EqualValues(t, "The quick brown fox.txt", link.Name)
	assert.EqualValues(t, 12345, link.Size)
	assert.EqualValues(t, "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12", link.Sha1)
	assert.EqualValues(t, "plain/text", link.ContentType)
	assert.EqualValues(t, "2011-01-26 13:33:37", link.UploadAt)
	assert.EqualValues(t, "https://abvzps.example.com/dl/l/4spxX_-cSO4/The+quick+brown+fox.txt", link.URL)
	assert.EqualValues(t, "4spxX_-cSO4", link.Token)
}

func TestFileInfo(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/info").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"72fA-_Lq8Ak6":{"id":"72fA-_Lq8Ak6","status":451,"name":"The quick brown fox.txt","size":123456789012,"sha1":"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12","content_type":"plain/text"}}}`)

	info, err := c().FileInfo("72fA-_Lq8Ak6")

	assert.Nil(t, err)
	assert.EqualValues(t, "72fA-_Lq8Ak6", info.ID)
	assert.EqualValues(t, 451, info.Status)
	assert.EqualValues(t, "The quick brown fox.txt", info.Name)
	assert.EqualValues(t, 123456789012, info.Size)
	assert.EqualValues(t, "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12", info.Sha1)
	assert.EqualValues(t, "plain/text", info.ContentType)
}

func TestFilesInfo(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/info").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"72fA-_Lq8Ak3":{"id":"72fA-_Lq8Ak3","status":200,"name":"The quick brown fox.txt","size":123456789012,"sha1":"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12","content_type":"plain/text"},"72fA-_Lq8Ak4":{"id":"72fA-_Lq8Ak4","status":500,"name":"The quick brown fox.txt","size":false,"sha1":"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12","content_type":"plain/text"},"72fA-_Lq8Ak5":{"id":"72fA-_Lq8Ak5","status":404,"name":false,"size":false,"sha1":false,"content_type":false},"72fA-_Lq8Ak6":{"id":"72fA-_Lq8Ak6","status":451,"name":"The quick brown fox.txt","size":123456789012,"sha1":"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12","content_type":"plain/text"}}}`)

	fileIDs := []string{
		"72fA-_Lq8Ak3",
		"72fA-_Lq8Ak4",
		"72fA-_Lq8Ak5",
		"72fA-_Lq8Ak6",
	}
	infos, err := c().FilesInfo(fileIDs)

	assert.Nil(t, err)
	assert.Len(t, infos, len(fileIDs))
	assert.Contains(t, infos, "72fA-_Lq8Ak3")
	assert.Contains(t, infos, "72fA-_Lq8Ak4")
	assert.Contains(t, infos, "72fA-_Lq8Ak5")
	assert.Contains(t, infos, "72fA-_Lq8Ak6")
}

func TestUploadLink(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/ul").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"url":"https://13abc37.example.com/ul/fCgaPthr_ys","valid_until":"2015-01-09 00:02:50"}}`)

	uploadLink, err := c().UploadLink("1234", "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12", false)

	assert.Nil(t, err)
	assert.EqualValues(t, "https://13abc37.example.com/ul/fCgaPthr_ys", uploadLink.URL)
	assert.EqualValues(t, "2015-01-09 00:02:50", uploadLink.ValidUntil)
}

func TestRemoteUpload(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/remotedl/add").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"id":"12","folderid":"4248"}}`)

	remote, err := c().RemoteUpload("http://google.com/favicon.ico", "folderID")

	assert.Nil(t, err)
	assert.EqualValues(t, "12", remote.ID)
	assert.EqualValues(t, "4248", remote.Folderid)
}

func TestRemoteUploadStatus(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/remotedl/status").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"3":{"id":3,"remoteurl":"http://127.0.0.1/","status":"error","bytes_loaded":"162","bytes_total":"162","folderid":"4","added":"2015-02-17 18:58:11","last_update":"2015-02-19 18:07:45","extid":false,"url":false},"20":{"id":20,"remoteurl":"http://google.de/favicon.ico","status":"finished","bytes_loaded":"229","bytes_total":"229","folderid":"4248","added":"2015-02-21 09:03:47","last_update":"2015-02-21 09:04:04","extid":"ANAaeBZus-Q","url":"https://openload.co/f/ANAaeBZus-Q"},"22":{"id":22,"remoteurl":"http://proof.ovh.net/files/1Gio.dat","status":"downloading","bytes_loaded":"823997062","bytes_total":"1073741824","folderid":"4248","added":"2015-02-21 09:20:26","last_update":"2015-02-21 09:21:56","extid":false,"url":false},"24":{"id":24,"remoteurl":"http://proof.ovh.net/files/100Mio.dat","status":"new","bytes_loaded":null,"bytes_total":null,"folderid":"4248","added":"2015-02-21 09:20:26","last_update":"2015-02-21 09:20:26","extid":false,"url":false}}}`)

	remote, err := c().RemoteUploadStatus(5, "1234")

	assert.Nil(t, err)
	assert.Len(t, remote, 4)
	assert.Contains(t, remote, "24")
	assert.Contains(t, remote, "22")
	assert.Contains(t, remote, "20")
	assert.Contains(t, remote, "3")
}

func TestListFolder(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/listfolder").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":{"folders":[{"id":"5144","name":".videothumb"},{"id":"5792","name":".subtitles"},{"id":"6272","name":"test"},{"id":"6288","name":"video"},{"id":"6396","name":"images"},{"id":"6990","name":"tmp"}],"files":[{"name":"big_buck_bunny.mp4.mp4","sha1":"c6531f5ce9669d6547023d92aea4805b7c45d133","folderid":"4258","upload_at":"1419791256","status":"active","size":"5114011","content_type":"video/mp4","download_count":"48","cstatus":"ok","link":"https://openload.co/f/UPPjeAk--30/big_buck_bunny.mp4.mp4","linkextid":"UPPjeAk--30"},{"name":"Sintel.2010.1080p.mkv.mp4","sha1":"7ca6da73b4f0881bd8dca78e9059e2e6830acce6","folderid":"4258","upload_at":"1426534681","status":"active","size":"1116102098","content_type":"video/mp4","download_count":"37","cstatus":"ok","link":"https://openload.co/f/AYgHe95d1E4/Sintel.2010.1080p.mkv.mp4","linkextid":"AYgHe95d1E4"}]}}`)

	listed, err := c().ListFolder("5")

	assert.Nil(t, err)
	assert.Len(t, listed.Folders, 6)
	assert.Len(t, listed.Files, 2)
}

func TestRenameFolder(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/renamefolder").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":true}`)

	renamed, err := c().RenameFolder("5", "my new foldername")

	assert.Nil(t, err)
	assert.EqualValues(t, true, renamed)
}

func TestRenameFile(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/rename").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":true}`)

	renamed, err := c().RenameFile("UPPjeAk--30", "My File Backup_2017.zip")

	assert.Nil(t, err)
	assert.EqualValues(t, true, renamed)
}

func TestDeleteFile(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/delete").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":true}`)

	deleted, err := c().DeleteFile("UPPjeAk--30")

	assert.Nil(t, err)
	assert.EqualValues(t, true, deleted)
}

func TestConvertFile(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/convert").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":true}`)

	converted, err := c().ConvertFile("UPPjeAk--30")

	assert.Nil(t, err)
	assert.EqualValues(t, true, converted)
}

func TestRunningConversions(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/runningconverts").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":[{"name":"Geysir.AVI","id":"3565411","status":"pending","last_update":"2015-08-23 19:41:40","progress":0.32,"retries":"0","link":"https://openload.co/f/f02JFG293J8/Geysir.AVI","linkextid":"f02JFG293J8"}]}`)

	conversions, err := c().RunningConversions("5")

	assert.Nil(t, err)
	assert.Len(t, conversions, 1)
	assert.EqualValues(t, "Geysir.AVI", conversions[0].Name)
	assert.EqualValues(t, "3565411", conversions[0].ID)
	assert.EqualValues(t, "pending", conversions[0].Status)
	assert.EqualValues(t, "2015-08-23 19:41:40", conversions[0].LastUpdate)
	assert.EqualValues(t, 0.32, conversions[0].Progress)
	assert.EqualValues(t, "0", conversions[0].Retries)
	assert.EqualValues(t, "https://openload.co/f/f02JFG293J8/Geysir.AVI", conversions[0].Link)
	assert.EqualValues(t, "f02JFG293J8", conversions[0].Linkextid)
}

func TestSplashImage(t *testing.T) {
	defer gock.Off()

	gock.New(buildAPIURL()).
		Get("/file/getsplash").
		Reply(200).
		BodyString(`{"status":200,"msg":"OK","result":"https://openload.co/splash/AYgHe95d1E4/zt8uSEmk56s.jpg"}`)

	splash, err := c().SplashImage("AYgHe95d1E4")

	assert.Nil(t, err)
	assert.EqualValues(t, "https://openload.co/splash/AYgHe95d1E4/zt8uSEmk56s.jpg", splash)
}
