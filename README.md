# lostfilm

The `lostfilm` library is a golang client to get episode`s torrent files from [lostfilm.tv](https://www.lostfilm.tv/) website.

## Install 

`go get github.com/bass3t/lostfilm`

## Usage

### Creating client

Create a client for processing connection to website.

```go
lf, err := lostfilm.NewClient()
```

### Login to server

Before sending requests to server need be autorize on server.

```go
err := lf.Login(userLogin, userPassword, nil)
```

Sometimes server request enter the captcha. If you can process this case, you need use the callback function. Captcha callback function gets the byte slice with captcha image and must return numeric recognized string with code.

```go
cb := func(c []byte) string {
    text, _ := recognizeCapture(c)
    return text
}

err := lf.Login(userLogin, userPassword, cb)
```

### Get Serials

Getting all serials from server.

```go
serials, err := lf.GetAllSerials()
```

Getting serials may be specified with filters.

```go
f := filter.Create()
f.Genre.Add(genre.Family)

serials, err := lf.GetSerials(filter)
```

Getting serial with alias (name). Alias is may be get form url of main page for serial `https://www.lostfilm.tv/series/The_Serial_Alias`.

```go
serial, err := lf.GetSerialByAlias("The_Serial_Alias")
```

### Get episode links

Getting all episodes description for serial.

```go
seasons, err := lf.GetSerialSeasons(serial)
```

Getting links for episode.

```go
links, err := lf.GetEpisodeLinks(&episode)
```
