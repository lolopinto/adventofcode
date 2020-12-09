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
	parts := splitLength(line, "contain", 2)

	bagDesc := strings.TrimSpace(parts[0])
	contents := strings.Split(parts[1], ",")

	b, ok := bagMap[bagDesc]
	// first time seeing it, add to map
	if !ok {
		b = &bag{
			desc: bagDesc,
		}
		bagMap[bagDesc] = b
	}

	for _, content := range contents {
		// nothing to do here
		content = strings.TrimSpace(content)
		content = strings.TrimRight(content, ".")
		if content == "no other bags" {
			continue
		}
		// pluralize
		if !strings.HasSuffix(content, "bags") {
			content = content + "s"
		}
		parts := strings.SplitN(content, " ", 2)
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
			// bag doesn't exist, add it
			bagC = &bag{
				desc: desc,
			}
			bagMap[desc] = bagC
		}
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
	for _, container := range b.containers {
		searchBag(container, foundContainers, needle)
		searchContainers(container, foundContainers, container.desc)
	}
}

func countBags(b *bag, bagMap map[string]*bag) int {
	sum := 0
	for _, content := range b.content {
		//		content.num

		desc := content.bag.desc
		b2, ok := bagMap[desc]
		if !ok {
			log.Fatalf("couldn't find bag %s", desc)
		}
		sumB := countBags(b2, bagMap)
		sum = sum + content.num + (content.num * sumB)
		//content.
	}
	return sum
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
	log.Println(countBags(goldbag, bagMap))
}
