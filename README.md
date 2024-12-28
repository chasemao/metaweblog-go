# metaweblog-go

metaweblog server developed by go, fully tested with Office Word 2016, support publish blog with image.

![](./word.png)

## Usage

1. New config file `config.yaml`, for example

```yaml
blogURL: https://chasemao.com        # The URL of your blog.
blogTitle: Chase Mao's blog         # The title of your blog.
blogDir: ../hugoblog/content/article/ # The directory where blog articles are saved.
mediaDir: ../hugoblog/static/media/  # The directory where media files (images, videos) are stored.
mediaRelDirForBlogHtml: /media/      # The relative path to the media directory, used in the blog HTML for links.
userName: xxx                        # Your username for authentication.
password: yyy                        # Your password for authentication.
```

- **`blogDir`**: The directory where your blog articles are stored.
  
- **`mediaDir`**: This is where you store all media files such as images, videos, etc.

- **`mediaRelDirForBlogHtml`**: This is a relative path used for linking media files inside your blog’s HTML content.

Here’s how the directories might look in a Hugo blog setup:

``` txt
hugosite/
├── archetypes/
├── assets/
├── content/   
│   └── article/         <-- blog article
├── data/
├── i18n/
├── layouts/
├── static/
│   └── media/           <-- medias used in article
├── themes/
└── hugo.toml
```                  


- **`/path/to/hugosite/content/article/`**: This is where your blog articles are saved.
  
- **`/path/to/hugosite/static/media/`**: This is where you store all media files like images, videos, etc.

Let’s say you have a media file, like an image, saved at the following location:

- File path: `/path/to/hugosite/static/media/a.jpeg`

In Hugo, files in the `static` directory are served at the root level of your site. So, the image `/path/to/hugosite/static/media/a.jpeg` would be accessed by visiting `/media/a.jpeg` in the final blog HTML.

If your `mediaRelDirForBlogHtml` is set to `/media/`, this means that Hugo will link to your media files in the HTML as `/media/a.jpeg`.

2. Run metaweblog go server on localhost:1314

```bash
metawebloggo
```

3. Optional flag `-a`, choose the address where the server is listening.