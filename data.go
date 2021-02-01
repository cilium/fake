// Copyright (C) Isovalent, Inc. - All Rights Reserved.
//
// NOTICE: All information contained herein is, and remains the property of
// Isovalent Inc and its suppliers, if any. The intellectual and technical
// concepts contained herein are proprietary to Isovalent Inc and its suppliers
// and may be covered by U.S. and Foreign Patents, patents in process, and are
// protected by trade secret or copyright law.  Dissemination of this information
// or reproduction of this material is strictly forbidden unless prior written
// permission is obtained from Isovalent Inc.

package fake

const alphanum = "abcdefghijklmnopqrstuvwxyz0123456789"

var tiers = []string{
	"dev",
	"integration",
	"testing",
	"staging",
	"prod",
}

var apps = []string{
	"cert-manager",
	"cilium",
	"coredns",
	"cyberark",
	"drupal",
	"elasticsearch",
	"etcd",
	"gitlab",
	"hubble-relay",
	"influxdb",
	"jenkins",
	"joomla",
	"kube-apiserver",
	"kube-proxy",
	"kube-scheduler",
	"magento",
	"mariadb",
	"mediawiki",
	"myelin",
	"pelican",
	"rabbitmq",
	"robin",
	"sonarqube",
	"wordpress",
	"zookeeper",
}
var labels = []string{
	"io.cilium/app",
	"k8s-app",
	"pod-template-generation",
	"pod-template-hash",
	// from https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/
	// and https://kubernetes.io/docs/reference/labels-annotations-taints/
	"app.kubernetes.io/component",
	"app.kubernetes.io/instance",
	"app.kubernetes.io/managed-by",
	"app.kubernetes.io/name",
	"app.kubernetes.io/part-of",
	"app.kubernetes.io/version",
	"kubernetes.io/arch",
	"kubernetes.io/hostname",
	"kubernetes.io/os",
	"node.kubernetes.io/instance-type",
	"topology.kubernetes.io/region",
	"topology.kubernetes.io/zone",
}

var namespaces = []string{
	"default",
	"kube-node-lease",
	"kube-public",
	"kube-system",
	"local-path-storage",
}

// from https://gist.github.com/hugsy/8910dc78d208e40de42deb29e62df913
var adjectives = []string{
	"abandoned", "absolute", "adorable", "adventurous", "academic", "acceptable", "acclaimed", "accomplished", "accurate", "aching", "acidic", "acrobatic", "active", "actual", "adept", "admirable", "admired", "adolescent", "adorable", "adored", "advanced", "afraid", "affectionate", "aged", "aggravating", "aggressive", "agile", "agitated", "agonizing", "agreeable", "ajar", "alarmed", "alarming", "alert", "alienated", "alive", "all", "altruistic", "amazing", "ambitious", "ample", "amused", "amusing", "anchored", "ancient", "angelic", "angry", "anguished", "animated", "annual", "another", "antique", "anxious", "any", "apprehensive", "appropriate", "apt", "arctic", "arid", "aromatic", "artistic", "ashamed", "assured", "astonishing", "athletic", "attached", "attentive", "attractive", "austere", "authentic", "authorized", "automatic", "avaricious", "average", "aware", "awesome", "awful", "awkward", "babyish", "bad", "back", "baggy", "bare", "barren", "basic", "beautiful", "belated", "beloved", "beneficial", "better", "best", "bewitched", "big", "big-hearted", "biodegradable", "bite-sized", "bitter", "black", "black-and-white", "bland", "blank", "blaring", "bleak", "blind", "blissful", "blond", "blue", "blushing", "bogus", "boiling", "bold", "bony", "boring", "bossy", "both", "bouncy", "bountiful", "bowed", "brave", "breakable", "brief", "bright", "brilliant", "brisk", "broken", "bronze", "brown", "bruised", "bubbly", "bulky", "bumpy", "buoyant", "burdensome", "burly", "bustling", "busy", "buttery", "buzzing", "calculating", "calm", "candid", "canine", "capital", "carefree", "careful", "careless", "caring", "cautious", "cavernous", "celebrated", "charming", "cheap", "cheerful", "cheery", "chief", "chilly", "chubby", "circular", "classic", "clean", "clear", "clear-cut", "clever", "close", "closed", "cloudy", "clueless", "clumsy", "cluttered", "coarse", "cold", "colorful", "colorless", "colossal", "comfortable", "common", "compassionate", "competent", "complete", "complex", "complicated", "composed", "concerned", "concrete", "confused", "conscious", "considerate", "constant", "content", "conventional", "cooked", "cool", "cooperative", "coordinated", "corny", "corrupt", "costly", "courageous", "courteous", "crafty", "crazy", "creamy", "creative", "creepy", "criminal", "crisp", "critical", "crooked", "crowded", "cruel", "crushing", "cuddly", "cultivated", "cultured", "cumbersome", "curly", "curvy", "cute", "cylindrical", "damaged", "damp", "dangerous", "dapper", "daring", "darling", "dark", "dazzling", "dead", "deadly", "deafening", "dear", "dearest", "decent", "decimal", "decisive", "deep", "defenseless", "defensive", "defiant", "deficient", "definite", "definitive", "delayed", "delectable", "delicious", "delightful", "delirious", "demanding", "dense", "dental", "dependable", "dependent", "descriptive", "deserted", "detailed", "determined", "devoted", "different", "difficult", "digital", "diligent", "dim", "dimpled", "dimwitted", "direct", "disastrous", "discrete", "disfigured", "disgusting", "disloyal", "dismal", "distant", "downright", "dreary", "dirty", "disguised", "dishonest", "dismal", "distant", "distinct", "distorted", "dizzy", "dopey", "doting", "double", "downright", "drab", "drafty", "dramatic", "dreary", "droopy", "dry", "dual", "dull", "dutiful", "each", "eager", "earnest", "early", "easy", "easy-going", "ecstatic", "edible", "educated", "elaborate", "elastic", "elated", "elderly", "electric", "elegant", "elementary", "elliptical", "embarrassed", "embellished", "eminent", "emotional", "empty", "enchanted", "enchanting", "energetic", "enlightened", "enormous", "enraged", "entire", "envious", "equal", "equatorial", "essential", "esteemed", "ethical", "euphoric", "even", "evergreen", "everlasting", "every", "evil", "exalted", "excellent", "exemplary", "exhausted", "excitable", "excited", "exciting", "exotic", "expensive", "experienced", "expert", "extraneous", "extroverted", "extra-large", "extra-small", "fabulous", "failing", "faint", "fair", "faithful", "fake", "false", "familiar", "famous", "fancy", "fantastic", "far", "faraway", "far-flung", "far-off", "fast", "fat", "fatal", "fatherly", "favorable", "favorite", "fearful", "fearless", "feisty", "feline", "female", "feminine", "few", "fickle", "filthy", "fine", "finished", "firm", "first", "firsthand", "fitting", "fixed", "flaky", "flamboyant", "flashy", "flat", "flawed", "flawless", "flickering", "flimsy", "flippant", "flowery", "fluffy", "fluid", "flustered", "focused", "fond", "foolhardy", "foolish", "forceful", "forked", "formal", "forsaken", "forthright", "fortunate", "fragrant", "frail", "frank", "frayed", "free", "French", "fresh", "frequent", "friendly", "frightened", "frightening", "frigid", "frilly", "frizzy", "frivolous", "front", "frosty", "frozen", "frugal", "fruitful", "full", "fumbling", "functional", "funny", "fussy", "fuzzy", "gargantuan", "gaseous", "general", "generous", "gentle", "genuine", "giant", "giddy", "gigantic", "gifted", "giving", "glamorous", "glaring", "glass", "gleaming", "gleeful", "glistening", "glittering", "gloomy", "glorious", "glossy", "glum", "golden", "good", "good-natured", "gorgeous", "graceful", "gracious", "grand", "grandiose", "granular", "grateful", "grave", "gray", "great", "greedy", "green", "gregarious", "grim", "grimy", "gripping", "grizzled", "gross", "grotesque", "grouchy", "grounded", "growing", "growling", "grown", "grubby", "gruesome", "grumpy", "guilty", "gullible", "gummy", "hairy", "half", "handmade", "handsome", "handy", "happy", "happy-go-lucky", "hard", "hard-to-find", "harmful", "harmless", "harmonious", "harsh", "hasty", "hateful", "haunting", "healthy", "heartfelt", "hearty", "heavenly", "heavy", "hefty", "helpful", "helpless", "hidden", "hideous", "high", "high-level", "hilarious", "hoarse", "hollow", "homely", "honest", "honorable", "honored", "hopeful", "horrible", "hospitable", "hot", "huge", "humble", "humiliating", "humming", "humongous", "hungry", "hurtful", "husky", "icky", "icy", "ideal", "idealistic", "identical", "idle", "idiotic", "idolized", "ignorant", "ill", "illegal", "ill-fated", "ill-informed", "illiterate", "illustrious", "imaginary", "imaginative", "immaculate", "immaterial", "immediate", "immense", "impassioned", "impeccable", "impartial", "imperfect", "imperturbable", "impish", "impolite", "important", "impossible", "impractical", "impressionable", "impressive", "improbable", "impure", "inborn", "incomparable", "incompatible", "incomplete", "inconsequential", "incredible", "indelible", "inexperienced", "indolent", "infamous", "infantile", "infatuated", "inferior", "infinite", "informal", "innocent", "insecure", "insidious", "insignificant", "insistent", "instructive", "insubstantial", "intelligent", "intent", "intentional", "interesting", "internal", "international", "intrepid", "ironclad", "irresponsible", "irritating", "itchy", "jaded", "jagged", "jam-packed", "jaunty", "jealous", "jittery", "joint", "jolly", "jovial", "joyful", "joyous", "jubilant", "judicious", "juicy", "jumbo", "junior", "jumpy", "juvenile", "kaleidoscopic", "keen", "key", "kind", "kindhearted", "kindly", "klutzy", "knobby", "knotty", "knowledgeable", "knowing", "known", "kooky", "kosher", "lame", "lanky", "large", "last", "lasting", "late", "lavish", "lawful", "lazy", "leading", "lean", "leafy", "left", "legal", "legitimate", "light", "lighthearted", "likable", "likely", "limited", "limp", "limping", "linear", "lined", "liquid", "little", "live", "lively", "livid", "loathsome", "lone", "lonely", "long", "long-term", "loose", "lopsided", "lost", "loud", "lovable", "lovely", "loving", "low", "loyal", "lucky", "lumbering", "luminous", "lumpy", "lustrous", "luxurious", "mad", "made-up", "magnificent", "majestic", "major", "male", "mammoth", "married", "marvelous", "masculine", "massive", "mature", "meager", "mealy", "mean", "measly", "meaty", "medical", "mediocre", "medium", "meek", "mellow", "melodic", "memorable", "menacing", "merry", "messy", "metallic", "mild", "milky", "mindless", "miniature", "minor", "minty", "miserable", "miserly", "misguided", "misty", "mixed", "modern", "modest", "moist", "monstrous", "monthly", "monumental", "moral", "mortified", "motherly", "motionless", "mountainous", "muddy", "muffled", "multicolored", "mundane", "murky", "mushy", "musty", "muted", "mysterious", "naive", "narrow", "nasty", "natural", "naughty", "nautical", "near", "neat", "necessary", "needy", "negative", "neglected", "negligible", "neighboring", "nervous", "new", "next", "nice", "nifty", "nimble", "nippy", "nocturnal", "noisy", "nonstop", "normal", "notable", "noted", "noteworthy", "novel", "noxious", "numb", "nutritious", "nutty", "obedient", "obese", "oblong", "oily", "oblong", "obvious", "occasional", "odd", "oddball", "offbeat", "offensive", "official", "old", "old-fashioned", "only", "open", "optimal", "optimistic", "opulent", "orange", "orderly", "organic", "ornate", "ornery", "ordinary", "original", "other", "our", "outlying", "outgoing", "outlandish", "outrageous", "outstanding", "oval", "overcooked", "overdue", "overjoyed", "overlooked", "palatable", "pale", "paltry", "parallel", "parched", "partial", "passionate", "past", "pastel", "peaceful", "peppery", "perfect", "perfumed", "periodic", "perky", "personal", "pertinent", "pesky", "pessimistic", "petty", "phony", "physical", "piercing", "pink", "pitiful", "plain", "plaintive", "plastic", "playful", "pleasant", "pleased", "pleasing", "plump", "plush", "polished", "polite", "political", "pointed", "pointless", "poised", "poor", "popular", "portly", "posh", "positive", "possible", "potable", "powerful", "powerless", "practical", "precious", "present", "prestigious", "pretty", "precious", "previous", "pricey", "prickly", "primary", "prime", "pristine", "private", "prize", "probable", "productive", "profitable", "profuse", "proper", "proud", "prudent", "punctual", "pungent", "puny", "pure", "purple", "pushy", "putrid", "puzzled", "puzzling", "quaint", "qualified", "quarrelsome", "quarterly", "queasy", "querulous", "questionable", "quick", "quick-witted", "quiet", "quintessential", "quirky", "quixotic", "quizzical", "radiant", "ragged", "rapid", "rare", "rash", "raw", "recent", "reckless", "rectangular", "ready", "real", "realistic", "reasonable", "red", "reflecting", "regal", "regular", "reliable", "relieved", "remarkable", "remorseful", "remote", "repentant", "required", "respectful", "responsible", "repulsive", "revolving", "rewarding", "rich", "rigid", "right", "ringed", "ripe", "roasted", "robust", "rosy", "rotating", "rotten", "rough", "round", "rowdy", "royal", "rubbery", "rundown", "ruddy", "rude", "runny", "rural", "rusty", "sad", "safe", "salty", "same", "sandy", "sane", "sarcastic", "sardonic", "satisfied", "scaly", "scarce", "scared", "scary", "scented", "scholarly", "scientific", "scornful", "scratchy", "scrawny", "second", "secondary", "second-hand", "secret", "self-assured", "self-reliant", "selfish", "sentimental", "separate", "serene", "serious", "serpentine", "several", "severe", "shabby", "shadowy", "shady", "shallow", "shameful", "shameless", "sharp", "shimmering", "shiny", "shocked", "shocking", "shoddy", "short", "short-term", "showy", "shrill", "shy", "sick", "silent", "silky", "silly", "silver", "similar", "simple", "simplistic", "sinful", "single", "sizzling", "skeletal", "skinny", "sleepy", "slight", "slim", "slimy", "slippery", "slow", "slushy", "small", "smart", "smoggy", "smooth", "smug", "snappy", "snarling", "sneaky", "sniveling", "snoopy", "sociable", "soft", "soggy", "solid", "somber", "some", "spherical", "sophisticated", "sore", "sorrowful", "soulful", "soupy", "sour", "Spanish", "sparkling", "sparse", "specific", "spectacular", "speedy", "spicy", "spiffy", "spirited", "spiteful", "splendid", "spotless", "spotted", "spry", "square", "squeaky", "squiggly", "stable", "staid", "stained", "stale", "standard", "starchy", "stark", "starry", "steep", "sticky", "stiff", "stimulating", "stingy", "stormy", "straight", "strange", "steel", "strict", "strident", "striking", "striped", "strong", "studious", "stunning", "stupendous", "stupid", "sturdy", "stylish", "subdued", "submissive", "substantial", "subtle", "suburban", "sudden", "sugary", "sunny", "super", "superb", "superficial", "superior", "supportive", "sure-footed", "surprised", "suspicious", "svelte", "sweaty", "sweet", "sweltering", "swift", "sympathetic", "tall", "talkative", "tame", "tan", "tangible", "tart", "tasty", "tattered", "taut", "tedious", "teeming", "tempting", "tender", "tense", "tepid", "terrible", "terrific", "testy", "thankful", "that", "these", "thick", "thin", "third", "thirsty", "this", "thorough", "thorny", "those", "thoughtful", "threadbare", "thrifty", "thunderous", "tidy", "tight", "timely", "tinted", "tiny", "tired", "torn", "total", "tough", "traumatic", "treasured", "tremendous", "tragic", "trained", "tremendous", "triangular", "tricky", "trifling", "trim", "trivial", "troubled", "true", "trusting", "trustworthy", "trusty", "truthful", "tubby", "turbulent", "twin", "ugly", "ultimate", "unacceptable", "unaware", "uncomfortable", "uncommon", "unconscious", "understated", "unequaled", "uneven", "unfinished", "unfit", "unfolded", "unfortunate", "unhappy", "unhealthy", "uniform", "unimportant", "unique", "united", "unkempt", "unknown", "unlawful", "unlined", "unlucky", "unnatural", "unpleasant", "unrealistic", "unripe", "unruly", "unselfish", "unsightly", "unsteady", "unsung", "untidy", "untimely", "untried", "untrue", "unused", "unusual", "unwelcome", "unwieldy", "unwilling", "unwitting", "unwritten", "upbeat", "upright", "upset", "urban", "usable", "used", "useful", "useless", "utilized", "utter", "vacant", "vague", "vain", "valid", "valuable", "vapid", "variable", "vast", "velvety", "venerated", "vengeful", "verifiable", "vibrant", "vicious", "victorious", "vigilant", "vigorous", "villainous", "violet", "violent", "virtual", "virtuous", "visible", "vital", "vivacious", "vivid", "voluminous", "wan", "warlike", "warm", "warmhearted", "warped", "wary", "wasteful", "watchful", "waterlogged", "watery", "wavy", "wealthy", "weak", "weary", "webbed", "wee", "weekly", "weepy", "weighty", "weird", "welcome", "well-documented", "well-groomed", "well-informed", "well-lit", "well-made", "well-off", "well-to-do", "well-worn", "wet", "which", "whimsical", "whirlwind", "whispered", "white", "whole", "whopping", "wicked", "wide", "wide-eyed", "wiggly", "wild", "willing", "wilted", "winding", "windy", "winged", "wiry", "wise", "witty", "wobbly", "woeful", "wonderful", "wooden", "woozy", "wordy", "worldly", "worn", "worried", "worrisome", "worse", "worst", "worthless", "worthwhile", "worthy", "wrathful", "wretched", "writhing", "wrong", "wry", "yawning", "yearly", "yellow", "yellowish", "young", "youthful", "yummy", "zany", "zealous", "zesty", "zigzag",
}

// from https://gist.github.com/hugsy/8910dc78d208e40de42deb29e62df913
var nouns = []string{
	"ability", "abroad", "abuse", "access", "accident", "account", "act", "action", "active", "activity", "actor", "ad", "addition", "address", "administration", "adult", "advance", "advantage", "advertising", "advice", "affair", "affect", "afternoon", "age", "agency", "agent", "agreement", "air", "airline", "airport", "alarm", "alcohol", "alternative", "ambition", "amount", "analysis", "analyst", "anger", "angle", "animal", "annual", "answer", "anxiety", "anybody", "anything", "anywhere", "apartment", "appeal", "appearance", "apple", "application", "appointment", "area", "argument", "arm", "army", "arrival", "art", "article", "aside", "ask", "aspect", "assignment", "assist", "assistance", "assistant", "associate", "association", "assumption", "atmosphere", "attack", "attempt", "attention", "attitude", "audience", "author", "average", "award", "awareness", "baby", "back", "background", "bad", "bag", "bake", "balance", "ball", "band", "bank", "bar", "base", "baseball", "basis", "basket", "bat", "bath", "bathroom", "battle", "beach", "bear", "beat", "beautiful", "bed", "bedroom", "beer", "beginning", "being", "bell", "belt", "bench", "bend", "benefit", "bet", "beyond", "bicycle", "bid", "big", "bike", "bill", "bird", "birth", "birthday", "bit", "bite", "bitter", "black", "blame", "blank", "blind", "block", "blood", "blow", "blue", "board", "boat", "body", "bone", "bonus", "book", "boot", "border", "boss", "bother", "bottle", "bottom", "bowl", "box", "boy", "boyfriend", "brain", "branch", "brave", "bread", "break", "breakfast", "breast", "breath", "brick", "bridge", "brief", "brilliant", "broad", "brother", "brown", "brush", "buddy", "budget", "bug", "building", "bunch", "burn", "bus", "business", "button", "buy", "buyer", "cabinet", "cable", "cake", "calendar", "call", "calm", "camera", "camp", "campaign", "can", "cancel", "cancer", "candidate", "candle", "candy", "cap", "capital", "car", "card", "care", "career", "carpet", "carry", "case", "cash", "cat", "catch", "category", "cause", "celebration", "cell", "chain", "chair", "challenge", "champion", "championship", "chance", "change", "channel", "chapter", "character", "charge", "charity", "chart", "check", "cheek", "chemical", "chemistry", "chest", "chicken", "child", "childhood", "chip", "chocolate", "choice", "church", "cigarette", "city", "claim", "class", "classic", "classroom", "clerk", "click", "client", "climate", "clock", "closet", "clothes", "cloud", "club", "clue", "coach", "coast", "coat", "code", "coffee", "cold", "collar", "collection", "college", "combination", "combine", "comfort", "comfortable", "command", "comment", "commercial", "commission", "committee", "common", "communication", "community", "company", "comparison", "competition", "complaint", "complex", "computer", "concentrate", "concept", "concern", "concert", "conclusion", "condition", "conference", "confidence", "conflict", "confusion", "connection", "consequence", "consideration", "consist", "constant", "construction", "contact", "contest", "context", "contract", "contribution", "control", "conversation", "convert", "cook", "cookie", "copy", "corner", "cost", "count", "counter", "country", "county", "couple", "courage", "course", "court", "cousin", "cover", "cow", "crack", "craft", "crash", "crazy", "cream", "creative", "credit", "crew", "criticism", "cross", "cry", "culture", "cup", "currency", "current", "curve", "customer", "cut", "cycle", "dad", "damage", "dance", "dare", "dark", "data", "database", "date", "daughter", "day", "dead", "deal", "dealer", "dear", "death", "debate", "debt", "decision", "deep", "definition", "degree", "delay", "delivery", "demand", "department", "departure", "dependent", "deposit", "depression", "depth", "description", "design", "designer", "desire", "desk", "detail", "development", "device", "devil", "diamond", "diet", "difference", "difficulty", "dig", "dimension", "dinner", "direction", "director", "dirt", "disaster", "discipline", "discount", "discussion", "disease", "dish", "disk", "display", "distance", "distribution", "district", "divide", "doctor", "document", "dog", "door", "dot", "double", "doubt", "draft", "drag", "drama", "draw", "drawer", "drawing", "dream", "dress", "drink", "drive", "driver", "drop", "drunk", "due", "dump", "dust", "duty", "ear", "earth", "ease", "east", "eat", "economics", "economy", "edge", "editor", "education", "effect", "effective", "efficiency", "effort", "egg", "election", "elevator", "emergency", "emotion", "emphasis", "employ", "employee", "employer", "employment", "end", "energy", "engine", "engineer", "engineering", "entertainment", "enthusiasm", "entrance", "entry", "environment", "equal", "equipment", "equivalent", "error", "escape", "essay", "establishment", "estate", "estimate", "evening", "event", "evidence", "exam", "examination", "example", "exchange", "excitement", "excuse", "exercise", "exit", "experience", "expert", "explanation", "expression", "extension", "extent", "external", "extreme", "eye", "face", "fact", "factor", "fail", "failure", "fall", "familiar", "family", "fan", "farm", "farmer", "fat", "father", "fault", "fear", "feature", "fee", "feed", "feedback", "feel", "feeling", "female", "few", "field", "fight", "figure", "file", "fill", "film", "final", "finance", "finding", "finger", "finish", "fire", "fish", "fishing", "fix", "flight", "floor", "flow", "flower", "fly", "focus", "fold", "following", "food", "foot", "football", "force", "forever", "form", "formal", "fortune", "foundation", "frame", "freedom", "friend", "friendship", "front", "fruit", "fuel", "fun", "function", "funeral", "funny", "future", "gain", "game", "gap", "garage", "garbage", "garden", "gas", "gate", "gather", "gear", "gene", "general", "gift", "girl", "girlfriend", "give", "glad", "glass", "glove", "go", "goal", "god", "gold", "golf", "good", "government", "grab", "grade", "grand", "grandfather", "grandmother", "grass", "great", "green", "grocery", "ground", "group", "growth", "guarantee", "guard", "guess", "guest", "guidance", "guide", "guitar", "guy", "habit", "hair", "half", "hall", "hand", "handle", "hang", "harm", "hat", "hate", "head", "health", "hearing", "heart", "heat", "heavy", "height", "hell", "hello", "help", "hide", "high", "highlight", "highway", "hire", "historian", "history", "hit", "hold", "hole", "holiday", "home", "homework", "honey", "hook", "hope", "horror", "horse", "hospital", "host", "hotel", "hour", "house", "housing", "human", "hunt", "hurry", "hurt", "husband", "ice", "idea", "ideal", "if", "illegal", "image", "imagination", "impact", "implement", "importance", "impress", "impression", "improvement", "incident", "income", "increase", "independence", "independent", "indication", "individual", "industry", "inevitable", "inflation", "influence", "information", "initial", "initiative", "injury", "insect", "inside", "inspection", "inspector", "instance", "instruction", "insurance", "intention", "interaction", "interest", "internal", "international", "internet", "interview", "introduction", "investment", "invite", "iron", "island", "issue", "it", "item", "jacket", "job", "join", "joint", "joke", "judge", "judgment", "juice", "jump", "junior", "jury", "keep", "key", "kick", "kid", "kill", "kind", "king", "kiss", "kitchen", "knee", "knife", "knowledge", "lab", "lack", "ladder", "lady", "lake", "land", "landscape", "language", "laugh", "law", "lawyer", "lay", "layer", "lead", "leader", "leadership", "leading", "league", "leather", "leave", "lecture", "leg", "length", "lesson", "let", "letter", "level", "library", "lie", "life", "lift", "light", "limit", "line", "link", "lip", "list", "listen", "literature", "living", "load", "loan", "local", "location", "lock", "log", "long", "look", "loss", "love", "low", "luck", "lunch", "machine", "magazine", "mail", "main", "maintenance", "major", "make", "male", "mall", "man", "management", "manager", "manner", "manufacturer", "many", "map", "march", "mark", "market", "marketing", "marriage", "master", "match", "mate", "material", "math", "matter", "maximum", "maybe", "meal", "meaning", "measurement", "meat", "media", "medicine", "medium", "meet", "meeting", "member", "membership", "memory", "mention", "menu", "mess", "message", "metal", "method", "middle", "midnight", "might", "milk", "mind", "mine", "minimum", "minor", "minute", "mirror", "miss", "mission", "mistake", "mix", "mixture", "mobile", "mode", "model", "mom", "moment", "money", "monitor", "month", "mood", "morning", "mortgage", "most", "mother", "motor", "mountain", "mouse", "mouth", "move", "movie", "mud", "muscle", "music", "nail", "name", "nasty", "nation", "national", "native", "natural", "nature", "neat", "necessary", "neck", "negative", "negotiation", "nerve", "net", "network", "news", "newspaper", "night", "nobody", "noise", "normal", "north", "nose", "note", "nothing", "notice", "novel", "number", "nurse", "object", "objective", "obligation", "occasion", "offer", "office", "officer", "official", "oil", "one", "opening", "operation", "opinion", "opportunity", "opposite", "option", "orange", "order", "ordinary", "organization", "original", "other", "outcome", "outside", "oven", "owner", "pace", "pack", "package", "page", "pain", "paint", "painting", "pair", "panic", "paper", "parent", "park", "parking", "part", "particular", "partner", "party", "pass", "passage", "passenger", "passion", "past", "path", "patience", "patient", "pattern", "pause", "pay", "payment", "peace", "peak", "pen", "penalty", "pension", "people", "percentage", "perception", "performance", "period", "permission", "permit", "person", "personal", "personality", "perspective", "phase", "philosophy", "phone", "photo", "phrase", "physical", "physics", "piano", "pick", "picture", "pie", "piece", "pin", "pipe", "pitch", "pizza", "place", "plan", "plane", "plant", "plastic", "plate", "platform", "play", "player", "pleasure", "plenty", "poem", "poet", "poetry", "point", "police", "policy", "politics", "pollution", "pool", "pop", "population", "position", "positive", "possession", "possibility", "possible", "post", "pot", "potato", "potential", "pound", "power", "practice", "preference", "preparation", "presence", "present", "presentation", "president", "press", "pressure", "price", "pride", "priest", "primary", "principle", "print", "prior", "priority", "private", "prize", "problem", "procedure", "process", "produce", "product", "profession", "professional", "professor", "profile", "profit", "program", "progress", "project", "promise", "promotion", "prompt", "proof", "property", "proposal", "protection", "psychology", "public", "pull", "punch", "purchase", "purple", "purpose", "push", "put", "quality", "quantity", "quarter", "queen", "question", "quiet", "quit", "quote", "race", "radio", "rain", "raise", "range", "rate", "ratio", "raw", "reach", "reaction", "read", "reading", "reality", "reason", "reception", "recipe", "recognition", "recommendation", "record", "recording", "recover", "red", "reference", "reflection", "refrigerator", "refuse", "region", "register", "regret", "regular", "relation", "relationship", "relative", "release", "relief", "remote", "remove", "rent", "repair", "repeat", "replacement", "reply", "report", "representative", "republic", "reputation", "request", "requirement", "research", "reserve", "resident", "resist", "resolution", "resolve", "resort", "resource", "respect", "respond", "response", "responsibility", "rest", "restaurant", "result", "return", "reveal", "revenue", "review", "revolution", "reward", "rice", "rich", "ride", "ring", "rip", "rise", "risk", "river", "road", "rock", "role", "roll", "roof", "room", "rope", "rough", "round", "routine", "row", "royal", "rub", "ruin", "rule", "run", "rush", "sad", "safe", "safety", "sail", "salad", "salary", "sale", "salt", "sample", "sand", "sandwich", "satisfaction", "save", "savings", "scale", "scene", "schedule", "scheme", "school", "science", "score", "scratch", "screen", "screw", "script", "sea", "search", "season", "seat", "second", "secret", "secretary", "section", "sector", "security", "selection", "self", "sell", "senior", "sense", "sensitive", "sentence", "series", "serve", "service", "session", "set", "setting", "sex", "shake", "shame", "shape", "share", "she", "shelter", "shift", "shine", "ship", "shirt", "shock", "shoe", "shoot", "shop", "shopping", "shot", "shoulder", "show", "shower", "sick", "side", "sign", "signal", "signature", "significance", "silly", "silver", "simple", "sing", "singer", "single", "sink", "sir", "sister", "site", "situation", "size", "skill", "skin", "skirt", "sky", "sleep", "slice", "slide", "slip", "smell", "smile", "smoke", "snow", "society", "sock", "soft", "software", "soil", "solid", "solution", "somewhere", "son", "song", "sort", "sound", "soup", "source", "south", "space", "spare", "speaker", "special", "specialist", "specific", "speech", "speed", "spell", "spend", "spirit", "spiritual", "spite", "split", "sport", "spot", "spray", "spread", "spring", "square", "stable", "staff", "stage", "stand", "standard", "star", "start", "state", "statement", "station", "status", "stay", "steak", "steal", "step", "stick", "still", "stock", "stomach", "stop", "storage", "store", "storm", "story", "strain", "stranger", "strategy", "street", "strength", "stress", "stretch", "strike", "string", "strip", "stroke", "structure", "struggle", "student", "studio", "study", "stuff", "stupid", "style", "subject", "substance", "success", "suck", "sugar", "suggestion", "suit", "summer", "sun", "supermarket", "support", "surgery", "surprise", "surround", "survey", "suspect", "sweet", "swim", "swimming", "swing", "switch", "sympathy", "system", "table", "tackle", "tale", "talk", "tank", "tap", "target", "task", "taste", "tax", "tea", "teach", "teacher", "teaching", "team", "tear", "technology", "telephone", "television", "tell", "temperature", "temporary", "tennis", "tension", "term", "test", "text", "thanks", "theme", "theory", "thing", "thought", "throat", "ticket", "tie", "till", "time", "tip", "title", "today", "toe", "tomorrow", "tone", "tongue", "tonight", "tool", "tooth", "top", "topic", "total", "touch", "tough", "tour", "tourist", "towel", "tower", "town", "track", "trade", "tradition", "traffic", "train", "trainer", "training", "transition", "transportation", "trash", "travel", "treat", "tree", "trick", "trip", "trouble", "truck", "trust", "truth", "try", "tune", "turn", "twist", "two", "type", "uncle", "understanding", "union", "unique", "unit", "university", "upper", "upstairs", "use", "user", "usual", "vacation", "valuable", "value", "variation", "variety", "vast", "vegetable", "vehicle", "version", "video", "view", "village", "virus", "visit", "visual", "voice", "volume", "wait", "wake", "walk", "wall", "war", "warning", "wash", "watch", "water", "wave", "way", "weakness", "wealth", "wear", "weather", "web", "wedding", "week", "weekend", "weight", "weird", "welcome", "west", "western", "wheel", "whereas", "while", "white", "whole", "wife", "will", "win", "wind", "window", "wine", "wing", "winner", "winter", "wish", "witness", "woman", "wonder", "wood", "word", "work", "worker", "working", "world", "worry", "worth", "wrap", "writer", "writing", "yard", "year", "yellow", "yesterday", "you", "young", "youth", "zone",
}
