package image

// Task is an image processing task
type Task struct {
	ImageID        string
	Width          int
	Height         int
	ApplyBlur      bool
	BlurAmount     int
	ApplyGrayscale bool
}

// NewTask creates a new image processing task
func NewTask(imageID string, width int, height int) *Task {
	return &Task{
		ImageID: imageID,
		Width:   width,
		Height:  height,
	}
}

// Blur applies gaussian blur to the image
func (t *Task) Blur(amount int) *Task {
	t.ApplyBlur = true
	t.BlurAmount = amount
	return t
}

// Grayscale turns the image into grayscale
func (t *Task) Grayscale() *Task {
	t.ApplyGrayscale = true
	return t
}
