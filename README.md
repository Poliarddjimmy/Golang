#Additional Information
##Features and Functions
Go’s file server has a few really nice features that are worth mentioning:
It sanitizes all request paths by running them through the path.Clean() function before
searching for a file. This removes any . and .. elements from the URL path, which helps
to stop directory traversal attacks. This feature is particularly useful if you’re using the
fileserver in conjunction with a router that doesn’t automatically sanitize URL paths.
Range requests are fully supported. This is great if your application is serving large files
and you want to support resumable downloads. You can see this functionality in action if
you use curl to request bytes 100-199 of the logo.png file, like so:
$ curl -i -H "Range: bytes=100-199" --output - http://localhost:4000/static/img/logo.png
HTTP/1.1 206 Partial Content
Accept-Ranges: bytes
Content-Length: 100
Content-Range: bytes 100-199/1075
Content-Type: image/png
Last-Modified: Thu, 04 May 2017 13:07:52 GMT
Date: Wed, 08 Aug 2018 16:21:16 GMT
[binary data]The Last-Modified and If-Modified-Since headers are transparently supported. If a
file hasn’t changed since the user last requested it, then http.FileServer will send a
304 Not Modified status code instead of the file itself. This helps reduce latency and
processing overhead for both the client and server.
The Content-Type is automatically set from the file extension using the
mime.TypeByExtension() function. You can add your own custom extensions and
content types using the mime.AddExtensionType() function if necessary.

##Performance
In the code above we’ve set up our file server so that it serves files out of the ./ui/static
directory on your hard disk.
But it’s important to note that, once the application is up-and-running, http.FileServer
probably won’t be reading these files from disk. Both Windows and Unix-based operating
systems cache recently-used files in RAM, so (for frequently-served files at least) it’s likely
that http.FileServer will be serving them from RAM rather than making the relatively slow
round-trip to your hard disk.

##Serving Single Files
Sometimes you might want to serve a single file from within a handler. For this there’s the
http.ServeFile() function, which you can use like so:
func downloadHandler(w http.ResponseWriter, r *http.Request) {
http.ServeFile(w, r, "./ui/static/file.zip")
}
Warning: http.ServeFile() does not automatically sanitize the file path. If you’re
constructing a file path from untrusted user input, to avoid directory traversal attacks
you must sanitize the input with filepath.Clean() before using it.
Disabling Directory Listings
If you want to disable directory listings there are a few different approaches you can take.
The simplest way? Add a blank index.html file to the specific directory that you want to
disable listings for. This will then be served instead of the directory listing, and the user will
get a 200 OK response with no body. If you want to do this for all directories under
./ui/static you can use the command:$ find ./ui/static -type d -exec touch {}/index.html \;
A more complicated (but arguably better) solution is to create a custom implementation of
http.FileSystem , and have it return an os.ErrNotExist error for any directories. A full
explanation and sample code can be found in this blog post.