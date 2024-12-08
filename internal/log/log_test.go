package log

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogSplited(t *testing.T) {
	lc := &LogConfig{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             "Info",
		Format:            "json",
		OutputPath:        "logs",
		SplitErrorLog:     true,
		Rotate: RotateOptions{
			MaxSize:    10,
			MaxAge:     1,
			MaxBackups: 1,
			Compress:   false,
		},
	}
	Init(lc)
	removeDirRecursively(lc.OutputPath)

	for i := 0; i < 30; i++ {
		Error("批量写入", String("content", "庆历四年的春天，滕子京被降职到巴陵郡做太守。隔了一年，政治清明通达，人民安居和顺，各种"))
		Error("批量写入", String("content", "庆历四年的春天，滕子京被降职到巴陵郡做太守。隔了一年，政治清明通达，人民安居和顺，各种"))
		Info("批量写入", String("content", "庆历四年的春天，滕子京被降职到巴陵郡做太守。隔了一年，政治清明通达，人民安居和顺，各种荒废的事业都兴办起来了。于是重新修建岳阳楼，扩大它原有的规模，把唐代名家和当代人的诗赋刻在它上面。嘱托我写一篇文章来记述这件事情。我观看那巴陵郡的美好景色，全在洞庭湖上。衔接远山，吞没长江，流水浩浩荡荡，无边无际，一天里阴晴多变，气象千变万化。这就是岳阳楼的雄伟景象。前人的记述（已经）很详尽了。那么向北面通到巫峡，向南面直到潇水和湘水，降职的官吏和来往的诗人，大多在这里聚会，（他们)观赏自然景物而触发的感情大概会有所不同吧？像那阴雨连绵，接连几个月不放晴，寒风怒吼，浑浊的浪冲向天空；太阳和星星隐藏起光辉，山岳隐没了形体；商人和旅客（一译：行商和客商）不能通行，船桅倒下，船桨折断；傍晚天色昏暗，虎在长啸，猿在悲啼，（这时）登上这座楼，就会有一种离开国都、怀念家乡，担心人家说坏话、惧怕人家批评指责，满眼都是萧条的景象，感慨到了极点而悲伤的心情。到了春风和煦，阳光明媚的时候，湖面平静，没有惊涛骇浪，天色湖光相连，一片碧绿，广阔无际；沙洲上的鸥鸟，时而飞翔，时而停歇，美丽的鱼游来游去，岸上的香草和小洲上的兰花，草木茂盛，青翠欲滴。有时大片烟雾完全消散，皎洁的月光一泻千里，波动的光闪着金色，静静的月影像沉入水中的玉璧，渔夫的歌声在你唱我和地响起来，这种乐趣（真是）无穷无尽啊！（这时）登上这座楼，就会感到心胸开阔、心情愉快，光荣和屈辱一并忘了，端着酒杯，吹着微风，觉得喜气洋洋了。哎呀！我曾探求过古时仁人的心境，或者和这些人的行为两样的，为什么呢？（是由于）不因外物好坏，自己得失而或喜或悲。在朝廷上做官时，就为百姓担忧；不在朝廷做官而处在僻远的江湖中间就为国君忧虑。他进也忧虑，退也忧愁。既然这样，那么他们什么时候才会感到快乐呢？古仁人必定说：“先于天下人的忧去忧，晚于天下人的乐去乐。”呀。唉！如果没有这种人，我与谁一道归去呢？写于为庆历六年九月十五日。"))
	}

	Sync()

	count_access, err1 := countLinesInFile(lc.OutputPath + "/access.log")
	count_error, err2 := countLinesInFile(lc.OutputPath + "/error.log")
	assert.Equal(t, 30, count_access, "Should generate 30 lines of error log")
	assert.Equal(t, 60, count_error, "Should generate 60 lines of error log")
	assert.Nil(t, err1, "Fail to read access log file")
	assert.Nil(t, err2, "Fail to read error log file")

}

func TestLogNotSplited(t *testing.T) {
	lc := &LogConfig{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             "Info",
		Format:            "json",
		OutputPath:        "logs",
		SplitErrorLog:     false,
		Rotate: RotateOptions{
			MaxSize:    10,
			MaxAge:     1,
			MaxBackups: 1,
			Compress:   false,
		},
	}
	Init(lc)
	removeDirRecursively(lc.OutputPath)

	for i := 0; i < 30; i++ {
		Error("批量写入", String("content", "庆历四年的春天，滕子京被降职到巴陵郡做太守。隔了一年，政治清明通达，人民安居和顺，各种"))
		Error("批量写入", String("content", "庆历四年的春天，滕子京被降职到巴陵郡做太守。隔了一年，政治清明通达，人民安居和顺，各种"))
		Info("批量写入", String("content", "庆历四年的春天，滕子京被降职到巴陵郡做太守。隔了一年，政治清明通达，人民安居和顺，各种荒废的事业都兴办起来了。于是重新修建岳阳楼，扩大它原有的规模，把唐代名家和当代人的诗赋刻在它上面。嘱托我写一篇文章来记述这件事情。我观看那巴陵郡的美好景色，全在洞庭湖上。衔接远山，吞没长江，流水浩浩荡荡，无边无际，一天里阴晴多变，气象千变万化。这就是岳阳楼的雄伟景象。前人的记述（已经）很详尽了。那么向北面通到巫峡，向南面直到潇水和湘水，降职的官吏和来往的诗人，大多在这里聚会，（他们)观赏自然景物而触发的感情大概会有所不同吧？像那阴雨连绵，接连几个月不放晴，寒风怒吼，浑浊的浪冲向天空；太阳和星星隐藏起光辉，山岳隐没了形体；商人和旅客（一译：行商和客商）不能通行，船桅倒下，船桨折断；傍晚天色昏暗，虎在长啸，猿在悲啼，（这时）登上这座楼，就会有一种离开国都、怀念家乡，担心人家说坏话、惧怕人家批评指责，满眼都是萧条的景象，感慨到了极点而悲伤的心情。到了春风和煦，阳光明媚的时候，湖面平静，没有惊涛骇浪，天色湖光相连，一片碧绿，广阔无际；沙洲上的鸥鸟，时而飞翔，时而停歇，美丽的鱼游来游去，岸上的香草和小洲上的兰花，草木茂盛，青翠欲滴。有时大片烟雾完全消散，皎洁的月光一泻千里，波动的光闪着金色，静静的月影像沉入水中的玉璧，渔夫的歌声在你唱我和地响起来，这种乐趣（真是）无穷无尽啊！（这时）登上这座楼，就会感到心胸开阔、心情愉快，光荣和屈辱一并忘了，端着酒杯，吹着微风，觉得喜气洋洋了。哎呀！我曾探求过古时仁人的心境，或者和这些人的行为两样的，为什么呢？（是由于）不因外物好坏，自己得失而或喜或悲。在朝廷上做官时，就为百姓担忧；不在朝廷做官而处在僻远的江湖中间就为国君忧虑。他进也忧虑，退也忧愁。既然这样，那么他们什么时候才会感到快乐呢？古仁人必定说：“先于天下人的忧去忧，晚于天下人的乐去乐。”呀。唉！如果没有这种人，我与谁一道归去呢？写于为庆历六年九月十五日。"))
	}

	Sync()

	count, err := countLinesInFile(lc.OutputPath + "/app.log")
	assert.Equal(t, 90, count, "Should generate 90 lines of log")
	assert.Nil(t, err, "Fail to read access log file")

}

func removeDirRecursively(path string) error {
	// 首先检查路径是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("路径 %s 不存在", path)
	}

	// 遍历目录中的所有文件和子目录
	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果有错误，返回
		}

		// 如果当前路径是目录，且不是根目录
		if info.IsDir() && currentPath != path {
			// 删除子目录
			err = os.Remove(currentPath)
			if err != nil {
				return fmt.Errorf("删除子目录 %s 时出错: %v", currentPath, err)
			}
		} else if !info.IsDir() {
			// 如果是文件，删除文件
			err = os.Remove(currentPath)
			if err != nil {
				return fmt.Errorf("删除文件 %s 时出错: %v", currentPath, err)
			}
		}

		return nil // 继续遍历
	})
	if err != nil {
		return fmt.Errorf("遍历或删除目录 %s 时出错: %v", path, err)
	}

	// 最后，删除根目录本身
	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("删除根目录 %s 时出错: %v", path, err)
	}

	return nil
}

func countLinesInFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lineCount := 0
	for {
		_, isPrefix, readErr := reader.ReadLine()
		if readErr == io.EOF {
			break // 文件结束
		}
		if readErr != nil {
			return 0, fmt.Errorf("读取文件时出错: %v", readErr)
		}

		if isPrefix {
			return 0, fmt.Errorf("文件过大，无法处理")
		}

		lineCount++
	}

	return lineCount, nil
}
