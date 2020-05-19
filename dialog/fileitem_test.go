package dialog

import (
	"path/filepath"
	"testing"

	"fyne.io/fyne/canvas"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"

	"github.com/stretchr/testify/assert"
)

func TestNewFileItem(t *testing.T) {
	f := &fileDialog{}
	_ = f.makeUI()
	item := f.newFileItem(canvas.NewImageFromResource(theme.FileIcon()), "/path/to/filename.txt", "filename", false)

	assert.Equal(t, item.name, "filename")

	test.Tap(item)
	assert.True(t, item.isCurrent)
	assert.Equal(t, item, f.selected)
}

func TestNewFileItem_Folder(t *testing.T) {
	f := &fileDialog{}
	_ = f.makeUI()
	currentDir, _ := filepath.Abs(".")
	parentDir := filepath.Dir(currentDir)
	f.setDirectory(parentDir)
	item := f.newFileItem(canvas.NewImageFromResource(theme.FolderIcon()), currentDir, filepath.Base(currentDir), true)

	assert.Equal(t, item.name, filepath.Base(currentDir))

	test.Tap(item)
	assert.False(t, item.isCurrent)
	assert.Equal(t, (*fileDialogItem)(nil), f.selected)
	assert.Equal(t, currentDir, f.dir)
}

func TestNewFileItem_ParentFolder(t *testing.T) {
	f := &fileDialog{}
	_ = f.makeUI()
	currentDir, _ := filepath.Abs(".")
	parentDir := filepath.Dir(currentDir)
	f.setDirectory(currentDir)
	item := f.newFileItem(canvas.NewImageFromResource(theme.FolderOpenIcon()), parentDir, "(Parent)", true)

	assert.Equal(t, item.name, "(Parent)")

	test.Tap(item)
	assert.False(t, item.isCurrent)
	assert.Equal(t, (*fileDialogItem)(nil), f.selected)
	assert.Equal(t, parentDir, f.dir)
}
