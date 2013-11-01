sajari-convert
==============

A Golang library to convert PDF, DOC, DOCX, XML, HTML, RTF, etc to plain text

The compiled binary runs as a service on port 8888 by default. Documents can be sent as a multipart POST request and the plain text (body) and meta information are then returned as a JSON object

This runs as a separate service as "os.Exec" causes a memory fork, so it's much more efficient to run this as a service with very low memory overhead.


