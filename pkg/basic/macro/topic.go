package macro

type Member struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Created int64  `json:"created"`
}

type Interactive struct {
	CommentNum int64 `json:"comment_num"`
	GoodNum    int64 `json:"good_num"`
	ShareNum   int64 `json:"share_num"`
}

type Topic struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Img         string      `json:"img"`
	Content     string      `json:"content"`
	Created     int64       `json:"created"`
	Member      Member      `json:"member"`
	Interactive Interactive `json:"interactive"`
}

var TopicData = []Topic{
	Topic{
		ID:      1,
		Name:    "新桥恋人",
		Img:     "https://pic1.zhimg.com/80/v2-2f2628bc63becb39824948932f796209_720w.jpeg",
		Content: "她一直被爱情抛弃，当终于有天遇到了真正的爱情，她却不能好好地看清他。为什么世间的爱情总是要经过如此多的磨难？",
		Created: 1609304637,
		Member: Member{
			ID:     351936818,
			Name:   "Deam Suresh",
			Avatar: "https://wx1.sinaimg.cn/mw690/66729e4cly1gmg85lrr9hj2050050dfq.jpg",
		},
		Interactive: Interactive{
			GoodNum:    216,
			CommentNum: 2,
			ShareNum:   390132,
		},
	},
	Topic{
		ID:      2,
		Name:    "恋情告急",
		Img:     "https://pic4.zhimg.com/80/v2-2fe26fdd713b1ad496122e06596d435e_720w.jpeg",
		Content: "现在我终于明白了，原来每一段恋爱，只要在心里面，已经是天长地久。",
		Created: 1609304637,
		Member: Member{
			ID:     351936818,
			Name:   "空山陋室",
			Avatar: "https://wx2.sinaimg.cn/mw690/66729e4cly1gmg85lrnlbj2050050q2s.jpg",
		},
		Interactive: Interactive{
			GoodNum:    16,
			CommentNum: 22,
			ShareNum:   132,
		},
	},
	Topic{
		ID:      3,
		Name:    "再说一次我爱你",
		Img:     "https://pic3.zhimg.com/80/v2-568edc5e812f6a436ab6c0e8d837d523_720w.jpeg",
		Content: "我真的很想念她，我现在才知道，原来想见一个人见不到的感觉是这样的。",
		Created: 1609304637,
		Member: Member{
			ID:     351936818,
			Name:   "Deam Suresh",
			Avatar: "https://tva2.sinaimg.cn/crop.0.0.179.179.180/66729e4cjw1ey7hz8k4lwj2050050q32.jpg",
		},
		Interactive: Interactive{
			GoodNum:    916,
			CommentNum: 403,
			ShareNum:   12765,
		},
	},
	Topic{
		ID:      4,
		Name:    "囧妈",
		Img:     "https://pic1.zhimg.com/80/v2-91a20757153e3c67c483481e07189c50_720w.jpeg",
		Content: "每个人都是独立的个体，每个个体都应该是完整的。爱不是控制和索取，爱是接纳和尊重，可惜我明白的太晚了",
		Created: 1609304637,
		Member: Member{
			ID:     351936818,
			Name:   "Deam Suresh",
			Avatar: "https://tvax2.sinaimg.cn/crop.0.0.200.200.180/78ed3187ly8gdi6uao6x9j205k05kq2t.jpg",
		},
		Interactive: Interactive{
			GoodNum:    89,
			CommentNum: 312,
			ShareNum:   67543,
		},
	},
	Topic{
		ID:      5,
		Name:    "野猪",
		Img:     "https://pic2.zhimg.com/80/v2-46c25c776511e0bc6e112d6bff2af277_720w.jpeg",
		Content: "我总是做着同一个梦，当所有我的伙伴儿们离开了家乡，我都会远远的看着他们。我知道，有一天你会从城市回来。",
		Created: 1609304637,
		Member: Member{
			ID:     351936818,
			Name:   "努力努力再努力",
			Avatar: "https://tvax4.sinaimg.cn/crop.0.0.1080.1080.180/a157f83bly8gfde69tn9aj20u00u03zp.jpg",
		},
		Interactive: Interactive{
			GoodNum:    90,
			CommentNum: 21,
			ShareNum:   20910,
		},
	},
}
