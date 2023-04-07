package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前项目的目录
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 获取项目 ID
	projectID, err := getProjectID(dir)
	if err != nil {
		panic(err)
	}
	if projectID == "" {
		panic("projectID empty")
	}
	// 备份书签
	if err := backupBookmarks(projectID); err != nil {
		panic(err)
	}

	fmt.Println("Bookmarks backup completed!")
}

func getProjectID(dir string) (string, error) {
	// 获取 workspace.xml 文件路径
	workspacePath := filepath.Join(dir, ".idea", "workspace.xml")

	// 读取 workspace.xml 文件内容
	content, err := os.ReadFile(workspacePath)
	if err != nil {
		return "", err
	}

	// 解析 workspace.xml 文件
	return getId(content), nil
}

type Project struct {
	XMLName    xml.Name    `xml:"project"`
	Components []Component `xml:"component"`
}
type Project2 struct {
	XMLName    xml.Name     `xml:"project"`
	Components []Component2 `xml:"component"`
}
type Component2 struct {
	XMLName xml.Name `xml:"component"`
	Name    string   `xml:"name,attr"`
	Content string   `xml:",innerxml"`
}
type Component struct {
	Name string `xml:"name,attr"`
	ID   string `xml:"id,attr"`
}

func getId(content []byte) string {

	//fmt.Println(string(content))
	// 解析 workspace.xml 文件
	var project Project
	if err := xml.Unmarshal(content, &project); err != nil {
		panic(err)
	}
	for _, component := range project.Components {
		if component.Name == "ProjectId" {
			return component.ID
		}
	}
	return ""
}
func bookmark(content []byte) []byte {
	var project Project2
	if err := xml.Unmarshal(content, &project); err != nil {
		panic(err)
	}
	for _, component := range project.Components {
		if component.Name == "BookmarksManager" {
			return []byte(component.Content)
		}
	}
	return []byte{}
}

func backupBookmarks(projectID string) error {
	// 获取 bookmarks 文件路径
	bookmarksPath := filepath.Join(
		os.Getenv("HOME"),
		"Library",
		"Application Support",
		"JetBrains",
		"GoLand2022.3",
		"workspace",
		projectID+".xml",
	)

	// 读取 bookmarks 文件内容
	content, err := os.ReadFile(bookmarksPath)
	if err != nil {
		return err
	}
	content = bookmark(content)
	// 将内容保存到备份文件中
	if err = os.WriteFile("bookmark-backup.xml", content, 0644); err != nil {
		return err
	}

	return nil
}
