package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/containers/storage/pkg/idtools"
	digest "github.com/opencontainers/go-digest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	wd := t.TempDir()

	pullOpts := map[string]string{"Test1": "test1", "Test2": "test2"}
	store, err := GetStore(StoreOptions{
		RunRoot:            filepath.Join(wd, "run"),
		GraphRoot:          filepath.Join(wd, "root"),
		GraphDriverName:    "vfs",
		GraphDriverOptions: []string{},
		UIDMap: []idtools.IDMap{{
			ContainerID: 0,
			HostID:      os.Getuid(),
			Size:        1,
		}},
		GIDMap: []idtools.IDMap{{
			ContainerID: 0,
			HostID:      os.Getgid(),
			Size:        1,
		}},
		PullOptions: pullOpts,
	})
	require.NoError(t, err)

	root := store.RunRoot()
	require.NotNil(t, root)

	root = store.GraphRoot()
	require.NotNil(t, root)

	root = store.GraphDriverName()
	require.NotNil(t, root)

	gopts := store.GraphOptions()
	assert.Equal(t, []string{}, gopts)

	store.UIDMap()
	store.GIDMap()

	opts := store.PullOptions()
	assert.Equal(t, pullOpts, opts)

	_, err = store.GraphDriver()
	require.Nil(t, err)

	_, err = store.CreateLayer("foo", "bar", nil, "", false, nil)
	require.Error(t, err)

	_, _, err = store.PutLayer("foo", "bar", nil, "", true, nil, nil)
	require.Error(t, err)

	_, err = store.CreateImage("foo", nil, "bar", "", nil)
	require.Error(t, err)

	_, err = store.CreateContainer("foo", nil, "bar", "layer", "", nil)
	require.Error(t, err)

	_, err = store.Metadata("foobar")
	require.Error(t, err)

	err = store.SetMetadata("foo", "bar")
	require.Error(t, err)

	exists := store.Exists("foobar")
	require.False(t, exists)

	_, err = store.Status()
	require.Nil(t, err)

	err = store.Delete("foobar")
	require.Error(t, err)

	err = store.DeleteLayer("foobar")
	require.Error(t, err)

	_, err = store.DeleteImage("foobar", true)
	require.Error(t, err)

	_, err = store.DeleteImage("foobar", false)
	require.Error(t, err)

	err = store.DeleteContainer("foobar")
	require.Error(t, err)

	err = store.DeleteContainer("foobar")
	require.Error(t, err)

	err = store.Wipe()
	require.Nil(t, err)

	_, err = store.Mount("foobar", "")
	require.Error(t, err)

	_, err = store.Unmount("foobar", true)
	require.Error(t, err)

	_, err = store.Unmount("foobar", false)
	require.Error(t, err)

	_, err = store.Mounted("foobar")
	require.Error(t, err)

	_, err = store.Changes("foobar", "foobar")
	require.Error(t, err)

	_, err = store.DiffSize("foobar", "foobar")
	require.Error(t, err)

	_, err = store.Diff("foobar", "foobar", nil)
	require.Error(t, err)

	_, err = store.ApplyDiff("foobar", nil)
	require.Error(t, err)

	var d digest.Digest
	_, err = store.LayersByCompressedDigest(d)
	require.Error(t, err)

	_, err = store.LayersByUncompressedDigest(d)
	require.Error(t, err)

	_, err = store.LayerSize("foobar")
	require.Error(t, err)

	_, _, err = store.LayerParentOwners("foobar")
	require.Error(t, err)

	_, err = store.Layers()
	require.Nil(t, err)

	_, err = store.Images()
	require.Nil(t, err)

	_, err = store.Containers()
	require.Nil(t, err)

	_, err = store.Names("foobar")
	require.Error(t, err)

	err = store.SetNames("foobar", nil)
	require.Error(t, err)

	_, err = store.ListImageBigData("foobar")
	require.Error(t, err)

	_, err = store.ImageBigData("foo", "bar")
	require.Error(t, err)

	_, err = store.ImageBigDataSize("foo", "bar")
	require.Error(t, err)

	_, err = store.ImageBigDataDigest("foo", "bar")
	require.Error(t, err)

	err = store.SetImageBigData("foo", "bar", nil, nil)
	require.Error(t, err)

	_, err = store.ImageSize("foobar")
	require.Error(t, err)

	_, err = store.ListContainerBigData("foobar")
	require.Error(t, err)

	_, err = store.ContainerBigData("foo", "bar")
	require.Error(t, err)

	_, err = store.ContainerBigDataSize("foo", "bar")
	require.Error(t, err)

	_, err = store.ContainerBigDataDigest("foo", "bar")
	require.Error(t, err)

	err = store.SetContainerBigData("foo", "bar", nil)
	require.Error(t, err)

	_, err = store.ContainerSize("foobar")
	require.Error(t, err)

	_, err = store.Layer("foobar")
	require.Error(t, err)

	_, err = store.Image("foobar")
	require.Error(t, err)

	_, err = store.ImagesByTopLayer("foobar")
	require.Error(t, err)

	images, err := store.ImagesByDigest("foobar")
	require.NoError(t, err)
	assert.Equal(t, len(images), 0)

	_, err = store.Container("foobar")
	require.Error(t, err)

	_, err = store.ContainerByLayer("foobar")
	require.Error(t, err)

	_, err = store.ContainerDirectory("foobar")
	require.Error(t, err)

	err = store.SetContainerDirectoryFile("foo", "bar", nil)
	require.Error(t, err)

	_, err = store.FromContainerDirectory("foo", "bar")
	require.Error(t, err)

	_, err = store.ContainerRunDirectory("foobar")
	require.Error(t, err)

	err = store.SetContainerRunDirectoryFile("foo", "bar", nil)
	require.Error(t, err)

	_, err = store.FromContainerRunDirectory("foo", "bar")
	require.Error(t, err)

	_, _, err = store.ContainerParentOwners("foobar")
	require.Error(t, err)

	_, err = store.Lookup("foobar")
	require.Error(t, err)

	_, err = store.Shutdown(false)
	require.Nil(t, err)

	_, err = store.Shutdown(true)
	require.Nil(t, err)

	_, err = store.Version()
	require.Nil(t, err)

	// GetDigestLock returns digest-specific Locker.
	_, err = store.GetDigestLock(d)
	require.Error(t, err)

	store.Free()
	store.Free()
}

func TestWithSplitStore(t *testing.T) {
	wd := t.TempDir()

	pullOpts := map[string]string{"Test1": "test1", "Test2": "test2"}
	store, err := GetStore(StoreOptions{
		RunRoot:            filepath.Join(wd, "run"),
		GraphRoot:          filepath.Join(wd, "root"),
		ImageStore:         filepath.Join(wd, "imgstore"),
		GraphDriverName:    "vfs",
		GraphDriverOptions: []string{},
		UIDMap: []idtools.IDMap{{
			ContainerID: 0,
			HostID:      os.Getuid(),
			Size:        1,
		}},
		GIDMap: []idtools.IDMap{{
			ContainerID: 0,
			HostID:      os.Getgid(),
			Size:        1,
		}},
		PullOptions: pullOpts,
	})
	require.NoError(t, err)

	root := store.RunRoot()
	require.NotNil(t, root)

	root = store.GraphRoot()
	require.NotNil(t, root)

	root = store.GraphDriverName()
	require.NotNil(t, root)

	gopts := store.GraphOptions()
	assert.Equal(t, []string{}, gopts)

	store.UIDMap()
	store.GIDMap()

	opts := store.PullOptions()
	assert.Equal(t, pullOpts, opts)

	_, err = store.GraphDriver()
	require.Nil(t, err)

	_, err = store.CreateLayer("foo", "bar", nil, "", false, nil)
	require.Error(t, err)

	_, _, err = store.PutLayer("foo", "bar", nil, "", true, nil, nil)
	require.Error(t, err)

	_, err = store.CreateImage("foo", nil, "bar", "", nil)
	require.Error(t, err)

	_, err = store.CreateContainer("foo", nil, "bar", "layer", "", nil)
	require.Error(t, err)

	_, err = store.Metadata("foobar")
	require.Error(t, err)

	err = store.SetMetadata("foo", "bar")
	require.Error(t, err)

	exists := store.Exists("foobar")
	require.False(t, exists)

	_, err = store.Status()
	require.Nil(t, err)

	err = store.Delete("foobar")
	require.Error(t, err)

	err = store.DeleteLayer("foobar")
	require.Error(t, err)

	_, err = store.DeleteImage("foobar", true)
	require.Error(t, err)

	_, err = store.DeleteImage("foobar", false)
	require.Error(t, err)

	err = store.DeleteContainer("foobar")
	require.Error(t, err)

	err = store.DeleteContainer("foobar")
	require.Error(t, err)

	err = store.Wipe()
	require.Nil(t, err)

	_, err = store.Mount("foobar", "")
	require.Error(t, err)

	_, err = store.Unmount("foobar", true)
	require.Error(t, err)

	_, err = store.Unmount("foobar", false)
	require.Error(t, err)

	_, err = store.Mounted("foobar")
	require.Error(t, err)

	_, err = store.Changes("foobar", "foobar")
	require.Error(t, err)

	_, err = store.DiffSize("foobar", "foobar")
	require.Error(t, err)

	_, err = store.Diff("foobar", "foobar", nil)
	require.Error(t, err)

	_, err = store.ApplyDiff("foobar", nil)
	require.Error(t, err)

	var d digest.Digest
	_, err = store.LayersByCompressedDigest(d)
	require.Error(t, err)

	_, err = store.LayersByUncompressedDigest(d)
	require.Error(t, err)

	_, err = store.LayerSize("foobar")
	require.Error(t, err)

	_, _, err = store.LayerParentOwners("foobar")
	require.Error(t, err)

	_, err = store.Layers()
	require.Nil(t, err)

	_, err = store.Images()
	require.Nil(t, err)

	_, err = store.Containers()
	require.Nil(t, err)

	_, err = store.Names("foobar")
	require.Error(t, err)

	err = store.SetNames("foobar", nil)
	require.Error(t, err)

	_, err = store.ListImageBigData("foobar")
	require.Error(t, err)

	_, err = store.ImageBigData("foo", "bar")
	require.Error(t, err)

	_, err = store.ImageBigDataSize("foo", "bar")
	require.Error(t, err)

	_, err = store.ImageBigDataDigest("foo", "bar")
	require.Error(t, err)

	err = store.SetImageBigData("foo", "bar", nil, nil)
	require.Error(t, err)

	_, err = store.ImageSize("foobar")
	require.Error(t, err)

	_, err = store.ListContainerBigData("foobar")
	require.Error(t, err)

	_, err = store.ContainerBigData("foo", "bar")
	require.Error(t, err)

	_, err = store.ContainerBigDataSize("foo", "bar")
	require.Error(t, err)

	_, err = store.ContainerBigDataDigest("foo", "bar")
	require.Error(t, err)

	err = store.SetContainerBigData("foo", "bar", nil)
	require.Error(t, err)

	_, err = store.ContainerSize("foobar")
	require.Error(t, err)

	_, err = store.Layer("foobar")
	require.Error(t, err)

	_, err = store.Image("foobar")
	require.Error(t, err)

	_, err = store.ImagesByTopLayer("foobar")
	require.Error(t, err)

	images, err := store.ImagesByDigest("foobar")
	require.NoError(t, err)
	assert.Equal(t, len(images), 0)

	_, err = store.Container("foobar")
	require.Error(t, err)

	_, err = store.ContainerByLayer("foobar")
	require.Error(t, err)

	_, err = store.ContainerDirectory("foobar")
	require.Error(t, err)

	err = store.SetContainerDirectoryFile("foo", "bar", nil)
	require.Error(t, err)

	_, err = store.FromContainerDirectory("foo", "bar")
	require.Error(t, err)

	_, err = store.ContainerRunDirectory("foobar")
	require.Error(t, err)

	err = store.SetContainerRunDirectoryFile("foo", "bar", nil)
	require.Error(t, err)

	_, err = store.FromContainerRunDirectory("foo", "bar")
	require.Error(t, err)

	_, _, err = store.ContainerParentOwners("foobar")
	require.Error(t, err)

	_, err = store.Lookup("foobar")
	require.Error(t, err)

	_, err = store.Shutdown(false)
	require.Nil(t, err)

	_, err = store.Shutdown(true)
	require.Nil(t, err)

	_, err = store.Version()
	require.Nil(t, err)

	// GetDigestLock returns digest-specific Locker.
	_, err = store.GetDigestLock(d)
	require.Error(t, err)

	store.Free()
	store.Free()
}
