package main

import (
	"log"
	"strconv"
	"strings"
)

type bagContent struct {
	num int
	bag *bag
}

type bag struct {
	desc       string
	content    []*bagContent
	containers []*bag
}

func parseLine(line string, bagMap map[string]*bag) {
	parts := strings.Split(line, "contain")
	if len(parts) != 2 {
		log.Fatalf("line %s is not as expected", line)
	}

	bagDesc := strings.TrimSpace(parts[0])
	contents := strings.Split(parts[1], ",")

	b, ok := bagMap[bagDesc]
	// first time seeing it, add to map
	if !ok {
		//		spew.Dump("creating " + bagDesc + " from source")
		b = &bag{
			desc: bagDesc,
		}
		bagMap[bagDesc] = b
	}

	//	spew.Dump(contents)
	for _, content := range contents {
		// nothing to do here
		content = strings.TrimSpace(content)
		content = strings.TrimRight(content, ".")
		//		spew.Dump(content)
		if content == "no other bags" {
			continue
		}
		// pluralize
		if !strings.HasSuffix(content, "bags") {
			content = content + "s"
		}
		parts := strings.SplitN(content, " ", 2)
		//		spew.Dump(parts)
		if len(parts) != 2 {
			log.Fatalf("content %s is not as expected", content)
		}
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		desc := parts[1]
		bagC, ok := bagMap[desc]
		if !ok {
			//			spew.Dump("creating " + desc + " from being contained")

			// bag doesn't exist, add it
			bagC = &bag{
				desc: desc,
			}
			bagMap[desc] = bagC
		}
		// bag already exists
		// it's a container
		// add it
		bagC.containers = append(bagC.containers, b)

		b.content = append(b.content, &bagContent{
			num: num,
			bag: bagC,
		})
	}
}

const toFind = "shiny gold bags"

func searchBag(b *bag, foundContainers map[string]bool, needle string) {
	for _, content := range b.content {
		if content.bag.desc == needle {
			foundContainers[b.desc] = true
		}
	}
}

func searchContainers(b *bag, foundContainers map[string]bool, needle string) {
	//	spew.Dump(b.desc)
	for _, container := range b.containers {
		searchBag(container, foundContainers, needle)
		//		searched, ok := searchedContainers[]
		// if container.desc == "muted yellow bags" {
		// 	spew.Dump(container)
		// }
		searchContainers(container, foundContainers, container.desc)
	}
}

func day7() {
	lines := readFile("day7input")
	bagMap := make(map[string]*bag)

	for _, line := range lines {
		parseLine(line, bagMap)
	}

	goldbag, ok := bagMap[toFind]

	if !ok {
		log.Fatal("couldn't find gold bag")
	}

	foundContainers := make(map[string]bool)
	//	searchedContainers := make(map[string]bool)

	searchContainers(goldbag, foundContainers, toFind)

	log.Println(len(foundContainers))
	//	spew.Dump(foundContainers)
}
