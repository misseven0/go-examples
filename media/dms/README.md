dms
dms is a UPnP DLNA Digital Media Server. It runs from the terminal, and serves content directly from the filesystem from the working directory, or the path given. The SSDP component will broadcast and respond to requests on all available network interfaces.

dms advertises and serves the raw files, in addition to alternate transcoded streams when it's able, such as mpeg2 PAL-DVD and WebM for the Chromecast. It will also provide thumbnails where possible.

dms also supports serving dynamic streams (e.g. a live rtsp stream) generated on the fly with the help of an external application (e.g. ffmpeg).

dms uses ffprobe/avprobe to get media data such as bitrate and duration, ffmpeg/avconv for video transoding, and ffmpegthumbnailer for generating thumbnails when browsing. These commands must be in the PATH given to dms or the features requiring them will be disabled.