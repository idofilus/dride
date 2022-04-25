@echo off

@REM CALL curl https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4 -o clips/BigBuckBunny.mp4
@REM CALL ffmpeg -i "clips/BigBuckBunny.mp4" -threads 3 -vcodec copy -f segment -segment_time 1:00 "clips/BigBuckBunny_h264_%02d.mp4"
@REM CALL for %%i in (clips/*.mp4) do ffmpeg -i "clips/%%i" -vf scale='min(150,iw)':-1 -f image2 -vframes 1 "clips/%%~ni.jpg"