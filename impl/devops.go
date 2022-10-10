package impl

import (
	"log"
	"os/exec"
)

func HexoCommandGenerate() error {
	log.Println("hexoCommand Generate")
	cmd := exec.Command("/home/pi/workdir/hexoDevOps/sh/pull_generate.sh")
	bytes, err := cmd.Output()
	if err != nil {
		log.Fatal("cmd output:", err)
		return err
	}
	log.Println(string(bytes))
	return nil
}

func HexoCommandNewPage(name string) error {
	log.Println("hexoCommand Newpage")
	cmd := exec.Command("hexo", "new", "post", "--cwd", "/home/pi/workdir/hexo/jinxycandotailwhip", name)
	bytes, err := cmd.Output()
	if err != nil {
		log.Fatal("cmd output:", err)
		return err
	}
	log.Println(string(bytes))
	cmd = exec.Command("./sh/push.sh")
	bytes, err = cmd.Output()
	if err != nil {
		log.Fatal("cmd output:", err)
		return err
	}
	log.Println(string(bytes))
	return nil
}
