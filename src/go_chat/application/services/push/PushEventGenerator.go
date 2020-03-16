package push

type PushEventGenerator struct {
	Events []PushObServer
}

func(p *PushEventGenerator) Add(server PushObServer){
	p.Events=append(p.Events,server)
}

func(p *PushEventGenerator) Update(){
	for _,event := range(p.Events){
		event.update();
	}
}