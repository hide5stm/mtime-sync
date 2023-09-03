package main

import (
	"os"
	"testing"
	"time"
)

func TestFindMaxSizeFile(t *testing.T) {
	// テスト用のディレクトリを作成
	testDir := "test_dir"
	err := os.Mkdir(testDir, os.ModeDir)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testDir)

	// テスト用のファイルを作成
	file1Path := testDir + "/file1.txt"
	file2Path := testDir + "/file2.txt"
	file3Path := testDir + "/file3.txt"

	// テスト用のファイル1 (最大サイズ)
	err = createTestFile(file1Path, 1024) // 1KB
	if err != nil {
		t.Fatal(err)
	}

	// テスト用のファイル2 (中間サイズ)
	err = createTestFile(file2Path, 512) // 0.5KB
	if err != nil {
		t.Fatal(err)
	}

	// テスト用のファイル3 (最小サイズ)
	err = createTestFile(file3Path, 256) // 0.25KB
	if err != nil {
		t.Fatal(err)
	}

	// findMaxSizeFileをテスト
	maxFilePath, maxFileModTime, err := FindMaxSizeFile(testDir)
	if err != nil {
		t.Fatal(err)
	}

	expectedMaxSize := int64(1024) // テスト用ファイル1のサイズ
	expectedModTime := getFileModTime(file1Path)

	// 最大サイズとmtimeを比較
	if maxFilePath != file1Path || maxFileModTime != expectedModTime {
		t.Errorf("findMaxSizeFile returned incorrect result. Expected: %s, %s; Got: %s, %s",
			file1Path, expectedModTime, maxFilePath, maxFileModTime)
	}
}

func createTestFile(filePath string, size int64) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 指定サイズのデータを書き込む
	data := make([]byte, size)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func getFileModTime(filePath string) time.Time {
	info, err := os.Stat(filePath)
	if err != nil {
		return time.Time{}
	}
	return info.ModTime()
}
