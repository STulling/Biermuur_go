package displaycontroller

var (
	Playlist []string = make([]string, 0)
	RequestNext chan [2]float64 = make(chan [2]float64, 0)
)

func requestNext(q audio.Queue) {
	rand.Seed(time.Now().Unix())
	streamer := io.Load(o.playlist[rand.Intn(len(o.playlist))])
	q.Add(streamer)
	q.Requested = false
}

func Orchestrate() {
	for {
        select {
        case <-RequestNext:
            requestNext(q audio.Queue)
        default:

        }
    }
}