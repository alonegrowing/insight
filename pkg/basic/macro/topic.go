package macro

import (
	"github.com/alonegrowing/insight/pkg/basic/dao"
)

type Topic = dao.Topic

var TopicData = []Topic{
	Topic{
		ID:      1,
		Name:    "新桥恋人",
		Img:     "https://pic1.zhimg.com/80/v2-2f2628bc63becb39824948932f796209_720w.jpeg",
		Content: "她一直被爱情抛弃，当终于有天遇到了真正的爱情，她却不能好好地看清他。为什么世间的爱情总是要经过如此多的磨难？",
		Created: 1609304637,
	},
	Topic{
		ID:      2,
		Name:    "恋情告急",
		Img:     "https://pic4.zhimg.com/80/v2-2fe26fdd713b1ad496122e06596d435e_720w.jpeg",
		Content: "现在我终于明白了，原来每一段恋爱，只要在心里面，已经是天长地久。",
		Created: 1609304637,
	},
	Topic{
		ID:      3,
		Name:    "再说一次我爱你",
		Img:     "https://pic3.zhimg.com/80/v2-568edc5e812f6a436ab6c0e8d837d523_720w.jpeg",
		Content: "我真的很想念她，我现在才知道，原来想见一个人见不到的感觉是这样的。",
		Created: 1609304637,
	},
	Topic{
		ID:      4,
		Name:    "囧妈",
		Img:     "https://pic1.zhimg.com/80/v2-91a20757153e3c67c483481e07189c50_720w.jpeg",
		Content: "每个人都是独立的个体，每个个体都应该是完整的。爱不是控制和索取，爱是接纳和尊重，可惜我明白的太晚了",
		Created: 1609304637,
	},
	Topic{
		ID:      5,
		Name:    "野猪",
		Img:     "https://pic2.zhimg.com/80/v2-46c25c776511e0bc6e112d6bff2af277_720w.jpeg",
		Content: "我总是做着同一个梦，当所有我的伙伴儿们离开了家乡，我都会远远的看着他们。我知道，有一天你会从城市回来。",
		Created: 1609304637,
	},
}
