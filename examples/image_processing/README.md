# Image Processing : Calculate phash on 1000 images

The idea is to download, decode and calculate phash on 1000 images, asynchronously. It handles error by itself, you won't see any `if err != nil` here :)

Each batch of size N is processed in parallel, and it uses a bit of everything from rinaugo Library.