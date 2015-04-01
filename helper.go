package wow

import (
    "fmt"
    "strconv"
    "strings"
    "html/template"
)

func GetItem(itemMap map[string]interface{}, slot string) *Item {
        item := itemMap[slot]
        switch item.(type) {
           case *Item:
                return item.(*Item)
           default:
                return nil
        }
}

func GetItemLink(item Item) map[string]template.HTML {
	var relParts []string

	if item.TooltipParams.Enchant > 0 {
		relParts = append(relParts, "&ench=", strconv.Itoa(item.TooltipParams.Enchant))
	}

        if item.TooltipParams.Tinker > 0 {
                relParts = append(relParts, "%ench=", strconv.Itoa(item.TooltipParams.Tinker))
        }


        if len(item.TooltipParams.Set) > 0 {
		set := make([]string, len(item.TooltipParams.Set))		
		for key, value := range item.TooltipParams.Set {
			set[key] = strconv.Itoa(value)
		}
                relParts = append(relParts, "&pcs=", strings.Join(set, ":"))
        }

	gems := make([]string, 0)
	if item.TooltipParams.Gem0 > 0 {
		gems = append(gems, strconv.Itoa(item.TooltipParams.Gem0))
	}
	if item.TooltipParams.Gem1 > 0 {
		gems = append(gems, strconv.Itoa(item.TooltipParams.Gem1))
	}
	if item.TooltipParams.Gem2 > 0 {
		gems = append(gems, strconv.Itoa(item.TooltipParams.Gem2))
	}

	if len(gems) > 0 {
                relParts = append(relParts, "&gems=", strings.Join(gems, ":"))
        }

	out := make(map[string]template.HTML)	

	out["url"] = template.HTML(fmt.Sprintf("https://www.wowhead.com/item=%s", strconv.Itoa(item.Id)))
	out["name"] = template.HTML(item.Name)
	out["rel"] = template.HTML(strings.Join(relParts, ""))

	return out
}
