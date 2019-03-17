# TorrentDeprivatiser
Replace the Announce URL with Public ones and remove the Private Bit on a folder full of .torrent

# Usage

You can find releases for various operating systems in the [releases tab](https://github.com/The-Eye-Team/TorrentDeprivatiser/releases).

Download one, then make it executable:

```
chmod +x TorrentDeprivatiser
```

Create a file with one announce url per line.
Sample usage with a folder called `torrents` with your torrent files inside:

```
./TorrentDeprivatiser -i torrents/ -t trackers.txt
```

You can see the options with the `-h` flag:

```
usage: TorrentDeprivatiser [-h|--help] -i|--input "<value>" [-j|--concurrency
                           <integer>] -t|--trackers "<value>"

                           Replace the Announce URL with Public ones and remove
                           the Private Bit on a folder full of .torrent

Arguments:

  -h  --help         Print help information
  -i  --input        Input directory
  -j  --concurrency  Concurrency. Default: 4
  -t  --trackers     Tracker list file
  ```
 
# Build

```
git clone https://github.com/The-Eye-Team/TorrentDeprivatiser.git && cd TorrentDeprivatiser
```

```
go get ./...
```

```
go build .
```

[![The-Eye.eu](https://the-eye.eu/public/.css/logo3_x300.png)](https://the-eye.eu)