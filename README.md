<h1 align="center"> 🥫 sauce - CLI anime identification

<h2 align="center"> sauce uses <a href="https://soruly.github.io/trace.moe-api/#/">trace.moe</a> to identify the anime in an image & serves you its essential details so you can determine if it's worth the watch. </h2>

<p align="center">
  <img src="https://github.com/cadecuddy/sauce/blob/main/resources/demo.gif" alt="demo" />
</p>

<h3 align="center"> <i>never find yourself asking <b>"sauce?"</b> ever again!</i>

#### This project was inspired by [what-anime-cli](https://github.com/irevenko/what-anime-cli/) by [irevenko](https://github.com/irevenko). I felt there was more utility in looking up the identified anime's MyAnimeList stats, so that was the approach I took with sauce. 

## 🔧 Installation

Install with [Go](https://go.dev/) install:
```bash
go install github.com/cadecuddy/sauce@latest
```

## 💻 Usage
**Known supported image types: jpg, png, webp, gif**

### 🔗 search by image url
`sauce url <url>`
```bash
sauce url https://findthis.jp/anime.png
```

### 📂 search by image file
`sauce file <path>`
```bash
sauce file demon-slayer.png
```

## 📟 Environment Setup (contributors)

`git clone https://github.com/cadecuddy/sauce.git` <br>
`cd sauce` <br>

If you find any bugs or want to add any cool features feel free to leave a PR!

## 🤝 made with
* [trace.moe](https://soruly.github.io/trace.moe-api/#/) - anime identification
* [jikan-go](https://github.com/darenliang/jikan-go) - MyAnimeList data
