# Videos CLI app
This app was built following [this tutorial](https://www.youtube.com/watch?v=CODqM_rzwtk&t=1320s). This repo is simply to showcase my learning journey. Please understand that this code is **not** my original work and that my projects may be written differently. Thank you!


## How to use this app
    # To generate an executable file, run:
    go build

### Get all videos

    # To get all videos from videos.json file, run:
    ./videos get --all
    
    # To get a single video from videos.json file, run:
    ./videos get --id {video id}



### Add a video

    # To add a video to the videos.json file, run:
    ./videos add --id {new id} --title {new title} --url {new url} --imageurl {new image url} --description {new description}