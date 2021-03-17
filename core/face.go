package core

type Face struct {
	FaceStatus FaceStatus
}

func (this *Face) SetStatus(s FaceStatus) {
	this.FaceStatus = s
}
