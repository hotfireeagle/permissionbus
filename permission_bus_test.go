package permissionbus

import (
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	filePath := filepath.Join(".", "asset", "example.json")
	filePath = filepath.Clean(filePath)

	pb, err := Load(filePath)

	if err != nil {
		t.Errorf(err.Error())
	}

	if pb.configData[0].Name != "权限管理" {
		t.Errorf("wrong parse")
	}

	if pb.configData[0].Children[0].Name != "角色管理" {
		t.Errorf("wrong parse")
	}

	if pb.configData[0].Children[0].Spec != menuType {
		t.Errorf("wrong parse")
	}

	if pb.configData[0].Children[0].Children[0].Name != "新增角色" {
		t.Errorf("wrong parse")
	}

	if pb.configData[0].Children[0].Children[1].Name != "获取角色列表" {
		t.Errorf("wrong parse")
	}
}

func TestRepeatScene(t *testing.T) {
	fp := filepath.Join(".", "asset", "repeat.json")
	fp = filepath.Clean(fp)

	_, err := Load(fp)
	if err == nil {
		t.Errorf("should accurs error")
	}
}

func TestApiCantHasChildren(t *testing.T) {
	fp := filepath.Join(".", "asset", "apiNoChildren.json")
	fp = filepath.Clean(fp)

	_, err := Load(fp)
	if err == nil {
		t.Errorf("should accurs error")
	}
}

func TestGetMenu(t *testing.T) {
	fp := filepath.Join(".", "asset", "example.json")
	fp = filepath.Clean(fp)

	ins, _ := Load(fp)

	menuTree := ins.GetMenuTree()

	if menuTree[0].Name != "权限管理" {
		t.Errorf("GetMenuTree方法出现问题")
	}

	if menuTree[1].Name != "数据管理" {
		t.Errorf("GetMenuTree方法出现问题")
	}

	if menuTree[0].Children[0].Name != "角色管理" {
		t.Errorf("GetMenuTree方法出现问题")
	}

	var check func(p PermissionConfigItem)
	check = func(p PermissionConfigItem) {
		if p.Spec == apiType {
			t.Errorf("在GetMenuTree中返回了api类型")
		}

		for _, i := range p.Children {
			check(i)
		}
	}

	for _, item := range menuTree {
		check(item)
	}
}

func TestApiGroupMustHasGroup(t *testing.T) {
	fp := filepath.Join(".", "asset", "errApiGroupNoGroup.json")
	fp = filepath.Clean(fp)

	_, err := Load(fp)
	if err == nil {
		t.Errorf("TestApiGroupMustHasGroup should accurs error")
	}
}

func TestApiGroupCantHasChild(t *testing.T) {
	fp := filepath.Join(".", "asset", "errApiGroupHasChildren.json")
	fp = filepath.Clean(fp)

	_, err := Load(fp)
	if err == nil {
		t.Errorf("TestApiGroupCantHasChild should accurs error")
	}
}

func TestApiGroupCantContainMenu(t *testing.T) {
	fp := filepath.Join(".", "asset", "errApiGroupHasMenuGroupChild.json")
	fp = filepath.Clean(fp)

	_, err := Load(fp)
	if err == nil {
		t.Errorf("TestApiGroupCantContainMenu should accurs error")
	}
}

func TestApiGroupCantContainApiGroup(t *testing.T) {
	fp := filepath.Join(".", "asset", "errApiGroupHasApiGroupGroupChild.json")
	fp = filepath.Clean(fp)

	_, err := Load(fp)
	if err == nil {
		t.Errorf("TestApiGroupCantContainApiGroup should accurs error")
	}
}

func TestMenuMustContainChildren(t *testing.T) {
	fp := filepath.Join(".", "asset", "errMenuNoChildren.json")
	fp = filepath.Clean(fp)

	_, err := Load(fp)
	if err == nil {
		t.Errorf("TestMenuMustContainChildren should accurs error")
	}
}

func TestMenuMustJustHaveMenuSibling(t *testing.T) {
	fp := filepath.Join(".", "asset", "errMenuMustJustHaveMenuSibling.json")
	fp = filepath.Clean(fp)

	_, err := Load(fp)
	if err == nil {
		t.Errorf("TestMenuMustJustHaveMenuSibling should accurs error")
	}
}

// 测试GetMenuByLeaf方法
func TestGetMenuByLeaf(t *testing.T) {
	fp := filepath.Join(".", "asset", "example.json")
	fp = filepath.Clean(fp)

	p, err := Load(fp)
	if err != nil {
		t.Errorf("TestGetMenuByLeaf should not return error")
	}

	selectedApis := []string{
		"编辑角色功能",
		"获取财务数据列表",
	}

	menus := p.GetMenuByLeaf(selectedApis)
	answer := []string{
		"权限管理",
		"角色管理",
		"数据管理",
		"财务数据",
	}

	if len(menus) != len(answer) {
		t.Errorf("error")
	}

	menuMap := make(map[string]bool)
	for _, m := range menus {
		menuMap[m] = true
	}

	for _, v := range answer {
		if menuMap[v] == false {
			t.Errorf("wrong")
		}
	}
}
