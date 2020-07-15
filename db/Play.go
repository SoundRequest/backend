package db

import (
	"fmt"
	"time"

	"github.com/SoundRequest/backend/structure"
)

func GetSongs(userid int) (data []structure.PlayItem, err error) {
	result := DB().Where("author = ?", userid).Find(&data)
	err = result.Error
	return
}

func AddSong(userid int, name, description, link string) error {
	result := DB().Create(&structure.PlayItem{Author: userid, Name: name, Description: description, Link: link})
	return result.Error
}

func UpdateSong(userid, target int, name, description, link string) error {
	result := DB().Model(&structure.PlayItem{}).Where("id = ? AND author = ?", target, userid).Update(&structure.PlayItem{Author: userid, Name: name, Description: description, Link: link, UpdatedAt: time.Now()})
	return result.Error
}

func RemoveSong(userid, target int) error {
	result := DB().Where("id = ? AND author = ?", target, userid).Delete(&structure.PlayItem{})
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	_ = DB().Where("item = ?", target).Delete(&structure.PlayBridge{})
	return nil
}

func GetList(userid int) (data []structure.PlayList, err error) {
	result := DB().Where("author = ?", userid).Find(&data)
	err = result.Error
	return
}

func GetListDetail(userid, listid int) ([]structure.PlayItem, error) {
	data := &structure.PlayList{}
	if _ = DB().Where("id = ?", listid).First(&data); data.Author != userid {
		return nil, ErrItemNotFound
	}
	itemData := &[]structure.PlayItem{}
	result := DB().Table("play_bridges as pb").Select("pi.*").Joins("JOIN play_items as pi on pi.id = pb.item").Where("play_list = ?", listid).Find(&itemData)
	return *itemData, result.Error
}

func AddList(userid int, name, description string, public bool) error {
	result := DB().Create(&structure.PlayList{Author: userid, Name: name, Description: description, Public: public})
	return result.Error
}

func UpdateList(userid, target int, name, description string, public bool) error {
	result := DB().Model(&structure.PlayList{}).Where("id = ? AND author = ?", target, userid).Update(&structure.PlayList{Author: userid, Name: name, Description: description, Public: public, UpdatedAt: time.Now()})
	return result.Error
}

func RemoveList(userid, target int) error {
	result := DB().Where("id = ? AND author = ?", target, userid).Delete(&structure.PlayList{})
	if result.Error != nil {
		return result.Error
	}

	_ = DB().Where("play_list = ?", target).Delete(&structure.PlayBridge{})
	return nil
}

func GetTag(userid int) (data []structure.PlayTag, err error) {
	result := DB().Where("author = ?", userid).Find(&data)
	err = result.Error
	return
}

func AddTag(userid int, name string) error {
	result := DB().Create(&structure.PlayTag{Author: userid, Name: name})
	return result.Error
}

func UpdateTag(userid, target int, name string) error {
	result := DB().Model(&structure.PlayTag{}).Where("id = ? AND author = ?", target, userid).Update(&structure.PlayList{Author: userid, Name: name, UpdatedAt: time.Now()})
	return result.Error
}

func RemoveTag(userid, target int) error {
	result := DB().Where("id = ? AND author = ?", target, userid).Delete(&structure.PlayTag{})
	if result.Error != nil {
		return result.Error
	}

	_ = DB().Where("play_tag = ?", target).Delete(&structure.PlayBridge{})
	return nil
}

func AddTagToItem(userid, itemid, tagid int) error {
	result := DB().Create(&structure.PlayBridge{Item: itemid, PlayTag: tagid})
	return result.Error
}

func RemoveTagFromItem(userid, itemid, tagid int) error {
	result := DB().Where("item = ? AND play_tag = ?", itemid, tagid).Delete(&structure.PlayBridge{})
	return result.Error
}

func AddPlayListToItem(userid, itemid, playlistid int) error {
	result := DB().Create(&structure.PlayBridge{Item: itemid, PlayList: playlistid})
	return result.Error
}

func RemovePlayListFromItem(userid, itemid, playlistid int) error {
	result := DB().Where("item = ? AND play_list = ?", itemid, playlistid).Delete(&structure.PlayBridge{})
	return result.Error
}
