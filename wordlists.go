package sentencecipher

// Word lists with 256 words each - Theme: Office/Workplace stories
// Each word list has 256 entries (0-255) to encode 1 full byte per word

var defaultNames = []string{
	// classic aliens
	"grey", "nordic", "reptiloid", "insectoid", "hybrid", "clone", "android", "synthoid",
	// star nations
	"pleiadian", "arcturian", "sirian", "orionite", "lyran", "andromedan", "venusian", "saturnian",
	// ancient astronauts
	"annunaki", "nephilim", "elohim", "watcher", "sentinel", "titan", "gibborim", "promethean",
	// gnostic horrors
	"archon", "demiurge", "yaldabaoth", "samael", "aeon", "logos", "sophia", "monad",
	// demonology
	"demon", "devil", "imp", "fiend", "incubus", "succubus", "cacodemon", "archfiend",
	// djinn & eastern spirits
	"djinn", "ifrit", "marid", "ghoul", "peri", "shaitan", "rakshasa", "asura",
	// angelology
	"angel", "seraph", "cherub", "archangel", "throne", "dominion", "principality", "ophanim",
	// ghosts
	"ghost", "phantom", "wraith", "specter", "revenant", "poltergeist", "banshee", "doppelganger",
	// undead
	"lich", "wight", "ghast", "shadowman", "fetch", "eidolon", "lemure", "phantasm",
	// elementals
	"sylph", "undine", "salamander", "gnome", "dryad", "nymph", "naiad", "oread",
	// cryptids & hollow earth
	"yeti", "mothman", "chupacabra", "agarthan", "hyperborean", "atlantean", "lemurian", "deros",
	// the agency
	"agent", "handler", "spook", "operative", "asset", "informant", "saboteur", "infiltrator",
	// occultists
	"occultist", "kabbalist", "gnostic", "theosophist", "alchemist", "astrologer", "diviner", "augur",
	// mages
	"sorcerer", "warlock", "magus", "thaumaturge", "theurge", "hierophant", "evocator", "necromancer",
	// psi-talents
	"psychic", "psion", "telepath", "empath", "precog", "clairvoyant", "medium", "channeler",
	// new age archetypes
	"lightworker", "starseed", "indigo", "wanderer", "wayshower", "gridkeeper", "earthkeeper", "gatekeeper",
	// gurus & hermits
	"guru", "swami", "yogi", "sage", "hermit", "mystic", "adept", "initiate",
	// prophets & oracles
	"prophet", "seer", "oracle", "sibyl", "haruspex", "vates", "pythia", "auspex",
	// thinking machines
	"mainframe", "singularity", "algorithm", "daemon", "golem", "cyborg", "replicant", "homunculus",
	// creepypasta
	"skinwalker", "wendigo", "rake", "crawler", "shade", "mimic", "husk", "hollow",
	// lovecraftiana
	"shoggoth", "nightgaunt", "byakhee", "elderthing", "deepone", "mi-go", "starspawn", "moonbeast",
	// unclassified horrors
	"horror", "abomination", "aberration", "anomaly", "cryptid", "chimera", "monstrosity", "atrocity",
	// shapeshifters
	"trickster", "shapeshifter", "werewolf", "vampire", "strigoi", "moroi", "upyr", "vrykolakas",
	// the fae
	"fae", "pixie", "sprite", "changeling", "goblin", "hobgoblin", "boggart", "redcap",
	// norse doomsday cast
	"valkyrie", "norn", "jotunn", "draugr", "fenrir", "nidhogg", "jormungandr", "sleipnir",
	// great old ones
	"cthulhu", "azathoth", "nyarlathotep", "yogsothoth", "shubniggurath", "hastur", "dagon", "ithaqua",
	// goetia & fallen kings
	"moloch", "baphomet", "mammon", "belial", "abaddon", "apollyon", "behemoth", "leviathan",
	// contactees
	"contactee", "abductee", "experiencer", "walkin", "dowser", "scryer", "charlatan", "mountebank",
	// cultists
	"cultist", "zealot", "fanatic", "acolyte", "neophyte", "templar", "druid", "shaman",
	// travelers between worlds
	"chrononaut", "dimensionaut", "planeswalker", "chronomancer", "dreamwalker", "oneironaut", "somnambulist", "sleeper",
	// psi-ops
	"psiop", "remoteviewer", "mentalist", "mentat", "esper", "sensitive", "psychonaut", "mesmerist",
	// the perps
	"gangstalker", "shill", "glowie", "puppeteer", "overlord", "wirepuller", "whisperer", "simulacrum",
}

// ==========================================
// THEME: BUSINESS / CORPORATE (Default)
// ==========================================

var businessSubjects = []string{
	"Project Update", "Meeting Notes", "Weekly Report", "Status Update",
	"Action Required", "Follow up", "Important Update", "Team Update",
	"Sync Up", "Discussion Points", "Next Steps", "Review Request",
	"Sprint Planning", "Budget Review", "Quarterly Goals", "Hiring Plan",
	"Client Feedback", "Marketing Strategy", "Sales Report", "Proposal Review",
}

var defaultVerbs = []string{
	// energy emission
	"radiates", "emanates", "vibrates", "resonates", "oscillates", "pulses", "hums", "thrums",
	// luminosity
	"glows", "shimmers", "flickers", "strobes", "flashes", "coruscates", "scintillates", "phosphoresces",
	// transmission
	"transmits", "broadcasts", "beams", "emits", "channels", "streams", "relays", "uplinks",
	// psionic ops
	"projects", "focuses", "amplifies", "attenuates", "dampens", "shields", "deflects", "grounds",
	// attunement
	"attunes", "aligns", "calibrates", "synchronizes", "entrains", "harmonizes", "modulates", "balances",
	// perception
	"senses", "intuits", "perceives", "envisions", "visualizes", "remoteviews", "scans", "dowses",
	// divination
	"foresees", "divines", "scries", "prophesies", "foretells", "predicts", "presages", "augurs",
	// altered states
	"retrocognizes", "psychometrizes", "regresses", "hypnotizes", "mesmerizes", "entrances", "trances", "sedates",
	// mind control
	"programs", "reprograms", "deprograms", "brainwashes", "indoctrinates", "gaslights", "manipulates", "influences",
	// domination
	"controls", "commands", "compels", "dominates", "subjugates", "enslaves", "possesses", "overrides",
	// summoning
	"invokes", "evokes", "summons", "conjures", "calls", "beckons", "entreats", "exhorts",
	// banishing
	"banishes", "exorcises", "dispels", "dissolves", "purges", "cleanses", "smudges", "nullifies",
	// malefica
	"hexes", "curses", "jinxes", "bewitches", "ensorcells", "glamours", "beguiles", "fascinates",
	// heavy cursing
	"execrates", "anathemizes", "excommunicates", "maledicts", "imprecates", "defiles", "desecrates", "blights",
	// consecration
	"consecrates", "anoints", "blesses", "sanctifies", "purifies", "hallows", "charges", "imprints",
	// manifestation & teleportation
	"manifests", "materializes", "dematerializes", "apparates", "apports", "bilocates", "teleports", "translocates",
	// levitation
	"levitates", "floats", "hovers", "glides", "drifts", "soars", "plummets", "ascends",
	// incarnation & haunting
	"descends", "incarnates", "reincarnates", "embodies", "haunts", "stalks", "infests", "pervades",
	// abduction protocol
	"abducts", "implants", "probes", "harvests", "tags", "clones", "splices", "mutates",
	// surveillance
	"tracks", "monitors", "surveils", "observes", "follows", "traces", "triangulates", "wiretaps",
	// gates & phase-shifting
	"portals", "gates", "phases", "blinks", "shifts", "folds", "warps", "unlocks",
	// tearing reality
	"tears", "rends", "sunders", "fractures", "splinters", "shatters", "collapses", "implodes",
	// quantum ops
	"entangles", "decoheres", "superposes", "tunnels", "leaps", "redshifts", "blueshifts", "spaghettifies",
	// time ops
	"rewinds", "loops", "echoes", "reverberates", "recurs", "dilates", "accelerates", "decelerates",
	// transmutation
	"transmutes", "metamorphoses", "transfigures", "evolves", "devolves", "morphs", "alchemizes", "sublimates",
	// awakening
	"expands", "contracts", "transcends", "awakens", "activates", "ignites", "triggers", "kindles",
	// merging
	"merges", "fuses", "integrates", "assimilates", "unifies", "coalesces", "interpenetrates", "commingles",
	// data transfer
	"downloads", "uploads", "encodes", "decodes", "encrypts", "decrypts", "ciphers", "transcribes",
	// revelation
	"unveils", "reveals", "discloses", "exposes", "leaks", "unmasks", "deciphers", "immanentizes",
	// voices from beyond
	"whispers", "murmurs", "mutters", "chants", "intones", "incants", "ululates", "keens",
	// psychic assault
	"zaps", "strikes", "smites", "lashes", "pierces", "penetrates", "breaches", "invades",
	// soul ops
	"reaps", "sows", "severs", "tethers", "anchors", "weaves", "unweaves", "unwinds",
}

var defaultObjects = []string{
	// psi-forms
	"thoughtforms", "tulpas", "egregores", "servitors", "constructs", "sigils", "glyphs", "wards",
	// talismanic
	"talismans", "amulets", "charms", "fetishes", "poppets", "effigies", "grisgris", "bindrunes",
	// subtle anatomy
	"chakras", "auras", "meridians", "nadis", "kundalini", "prana", "chi", "orgone",
	// exotic physics
	"torsionfields", "scalarwaves", "tachyons", "chronons", "gravitons", "plasmas", "antimatter", "darkmatter",
	// aetheric substances
	"aether", "vril", "blackgoo", "odyl", "ormus", "quintessence", "phlogiston", "miasma",
	// waves & vibes
	"frequencies", "vibrations", "oscillations", "waveforms", "wavelengths", "harmonics", "resonances", "biorhythms",
	// psychotronic hardware
	"psychotrons", "radionics", "accumulators", "montaukchairs", "dreamachines", "emitters", "cloudbusters", "teslacoils",
	// implant tech
	"implants", "microchips", "nanobots", "cybernetics", "neurochips", "electrodes", "brainwaves", "alphawaves",
	// entrainment media
	"thetawaves", "betawaves", "gammawaves", "binaurals", "isochronics", "subliminals", "affirmations", "mantras",
	// the seven planes
	"astralplane", "ethericplane", "mentalplane", "causalplane", "buddhicplane", "atmicplane", "voidplane", "shadowplane",
	// dimensional address space
	"dimensions", "densities", "timelines", "multiverses", "realities", "realms", "planes", "branes",
	// gates
	"portals", "gateways", "wormholes", "stargates", "doorways", "thresholds", "triangles", "leylines",
	// earth grid
	"grids", "nodes", "nexuses", "vortices", "earthchakras", "obelisks", "monoliths", "agartha",
	// megalithic
	"pyramids", "ziggurats", "megaliths", "menhirs", "barrows", "tumuli", "labyrinths", "sphinxes",
	// ritual tools
	"crystals", "geodes", "skulls", "cairns", "wands", "athames", "censers", "cauldrons",
	// divination
	"tarots", "runes", "pendulums", "rods", "dice", "tealeaves", "bones", "entrails",
	// forbidden texts
	"grimoires", "codices", "tomes", "necronomicons", "voynich", "palimpsests", "incunabula", "apocrypha",
	// souls & cords
	"souls", "spirits", "essences", "atmans", "jivas", "lightbodies", "silvercords", "shadowselves",
	// jungian inventory
	"egos", "ids", "animas", "animuses", "archetypes", "complexes", "personas", "daimones",
	// eastern metaphysics
	"karmas", "samskaras", "nirvanas", "samsara", "maya", "prakriti", "kaliyuga", "mandalas",
	// more subtle anatomy
	"thirdeyes", "pineals", "ajnas", "merkebas", "halos", "coronas", "sephirot", "qliphoth",
	// kabbalah & number mysticism
	"gematrias", "numerologies", "enneagrams", "tetragrammatons", "daaths", "ainsofs", "gnosis", "philosopherstones",
	// deep space
	"blackholes", "eventhorizons", "singularities", "quasars", "pulsars", "magnetars", "neutronstars", "supernovae",
	// the sky
	"galaxies", "nebulae", "constellations", "zodiacs", "asteroids", "comets", "planetx", "nibiru",
	// the craft
	"saucers", "motherships", "orbs", "boomerangs", "foofighters", "dieglocke", "scoutships", "skyfish",
	// abduction evidence
	"cropcircles", "scoopmarks", "mutilations", "abductions", "missingtime", "screenmemories", "hybridchildren", "starchildren",
	// the conspiracy shelf
	"chemtrails", "morgellons", "haarp", "gwentowers", "fluoride", "vaccines", "smartdust", "bluebeams",
	// black projects
	"mkultra", "montauk", "tavistock", "blackprojects", "skunkworks", "area51", "dulce", "hangar18",
	// radiation
	"gammarays", "xrays", "microwaves", "elfwaves", "emfs", "schumannwaves", "cosmicrays", "vanallenbelts",
	// the dream sector
	"dreams", "nightmares", "visions", "hallucinations", "delusions", "hypnagogia", "sleepparalysis", "luciddreams",
	// hyperspace pharmacology
	"machineelves", "hyperspace", "dmt", "ayahuasca", "mescaline", "psilocybin", "ketamine", "amrita",
	// the big stuff
	"akashicrecords", "noosphere", "morphicfields", "collectiveunconscious", "oversoul", "gaia", "dreamtime", "eschaton",
}

// ==========================================
// THEME: TECHNOLOGY / ENGINEERING
// ==========================================

var techSubjects = []string{
	"Server Outage", "Deployment Status", "API Updates", "Security Alert",
	"Code Review", "Database Migration", "System Maintenance", "Bug Bash",
	"Incident Report", "Release Notes", "Latency Issues", "Network Upgrade",
	"Cloud Infrastructure", "DevOps Sync", "Architecture Review", "Feature Flag",
}

// 256 Tech verbs
var techVerbs = []string{
	"codes", "programs", "compiles", "debugs", "executes", "runs", "builds", "deploys",
	"commits", "pushes", "pulls", "merges", "branches", "forks", "clones", "rebases",
	"fetches", "checkouts", "stashes", "reverts", "resets", "cherry-picks", "tags", "releases",
	"installs", "updates", "upgrades", "patches", "configures", "sets", "initializes", "boots",
	"restarts", "reboots", "shuts", "terminates", "kills", "stops", "halts", "suspends",
	"monitors", "logs", "tracks", "traces", "alerts", "notifies", "pings", "queries",
	"requests", "responses", "retrieves", "posts", "gets", "puts", "deletes", "corrects",
	"authenticates", "authorizes", "verifies", "validates", "encrypts", "decrypts", "hashes", "signs",
	"compresses", "zips", "archives", "extracts", "parses", "serializes", "encodes", "decodes",
	"renders", "paints", "draws", "displays", "shows", "hides", "toggles", "switches",
	"clicks", "taps", "scrolls", "swipes", "drags", "drops", "hovers", "focuses",
	"inputs", "types", "pastes", "copies", "cuts", "selects", "highlights", "edits",
	"saves", "loads", "reads", "writes", "opens", "closes", "imports", "exports",
	"uploads", "downloads", "syncs", "transfers", "streams", "buffers", "caches", "stores",
	"indexes", "searches", "filters", "sorts", "groups", "maps", "reduces", "iterates",
	"loops", "breaks", "continues", "returns", "throws", "catches", "tries", "awaits",
	"resolves", "rejects", "promises", "observes", "subscribes", "publishes", "emits", "broadcasts",
	"connects", "disconnects", "listens", "binds", "unbinds", "mounts", "unmounts", "routes",
	"navigates", "redirects", "forwards", "proxies", "tunnels", "bridges", "links", "chains",
	"tests", "mocks", "stubs", "spies", "asserts", "expects", "fails", "passes",
	"measures", "profiles", "benchmarks", "optimizes", "scales", "expands", "shrinks", "limits",
	"throttles", "blocks", "allows", "denies", "bans", "permits", "grants", "revokes",
	"assigns", "allocates", "frees", "collects", "cleans", "wipes", "erases", "formats",
	"partitions", "attaches", "ejects", "scans", "pairs", "charges", "powers", "drains",
	"computes", "calculates", "processes", "handles", "services", "providers", "consumes", "generates",
	"simulates", "emulates", "virtualizes", "containers", "dockers", "kubes", "orchestrates", "manages",
	"automates", "scripts", "commands", "controls", "rules", "governs", "directs", "leads",
	"analyzes", "inspects", "audits", "reviews", "checks", "examines", "probes", "assesses",
	"designs", "architects", "engineers", "develops", "implements", "integrates", "unifies", "combines",
	"refactors", "rewrites", "refines", "polishes", "improves", "fixes", "solves", "hacks",
	"tweaks", "tunes", "adjusts", "modifies", "changes", "transforms", "converts", "migrates",
	"deprecates", "removes", "scraps", "destroys", "discards", "purges", "prunes", "trims",
}

// 256 Tech objects
var techObjects = []string{
	"servers", "databases", "tables", "rows", "columns", "indexes", "views", "queries",
	"apis", "endpoints", "routes", "paths", "params", "headers", "bodies", "payloads",
	"tokens", "keys", "secrets", "passwords", "hashes", "salts", "ciphers", "certs",
	"logs", "metrics", "traces", "spans", "events", "signals", "alerts", "errors",
	"exceptions", "bugs", "issues", "tickets", "tasks", "stories", "epics", "sprints",
	"commits", "prs", "branches", "tags", "releases", "builds", "artifacts", "images",
	"containers", "pods", "nodes", "clusters", "services", "ingresses", "volumes", "networks",
	"firewalls", "rules", "policies", "roles", "groups", "users", "accounts", "profiles",
	"sessions", "cookies", "caches", "buffers", "queues", "topics", "channels", "streams",
	"sockets", "ports", "ips", "subnets", "dns", "records", "domains", "hosts",
	"files", "folders", "dirs", "locations", "links", "inodes", "blocks", "sectors",
	"disks", "drives", "ssds", "hdds", "raids", "partitions", "mounts", "swaps",
	"cpu", "ram", "gpu", "tpu", "cores", "threads", "processes", "daemons",
	"kernels", "drivers", "modules", "firmware", "bios", "bootloaders", "os", "distros",
	"shells", "terminals", "consoles", "scripts", "binaries", "libraries", "packages", "dependencies",
	"plugins", "imports", "exports", "classes", "objects", "functions", "methods", "variables",
	"constants", "types", "interfaces", "structs", "enums", "arrays", "lists", "maps",
	"sets", "trees", "graphs", "heaps", "stacks", "vectors", "matrices", "tensors",
	"strings", "integers", "floats", "booleans", "bytes", "bits", "chars", "runes",
	"pointers", "refs", "values", "scopes", "closures", "contexts", "promises", "futures",
	"fibers", "mutexes", "locks", "semaphores", "monitors", "conditions", "barriers", "latches",
	"frontend", "backend", "fullstack", "ui", "ux", "css", "html", "js",
	"react", "vue", "angular", "node", "deno", "bun", "go", "rust",
	"java", "python", "ruby", "php", "c", "cpp", "csharp", "swift",
	"kotlin", "scala", "clojure", "haskell", "elixir", "erlang", "lua", "perl",
	"sql", "nosql", "redis", "mongo", "postgres", "mysql", "sqlite", "oracle",
	"aws", "azure", "gcp", "cloud", "lambda", "s3", "ec2", "rds",
	"docker", "kubernetes", "helm", "terraform", "ansible", "jenkins", "gitlab", "github",
	"vscode", "vim", "emacs", "nano", "ide", "editor", "compiler", "debugger",
	"linter", "formatter", "parser", "lexer", "ast", "bytecode", "assembly", "machine",
	"laptop", "desktop", "monitor", "keyboard", "mouse", "trackpad", "webcam", "mic",
	"wifi", "bluetooth", "ethernet", "cable", "router", "switch", "modem", "gateway",
}

// Common components
var emailOpeners = []string{
	"Hi Team,",
	"Dear Colleagues,",
	"Hello everyone,",
	"Hi all,",
	"Good morning,",
	"Good afternoon,",
	"Team,",
	"Greetings,",
	"Hi there,",
	"All,",
}

// Connectors to chain sentences naturally
var sentenceConnectors = []string{
	"Please note that",
	"Additionally,",
	"Furthermore,",
	"Also,",
	"Moreover,",
	"In the meantime,",
	"Moving forward,",
	"As discussed,",
	"Just a reminder that",
	"For your information,",
	"However,",
	"Consequently,",
	"Therefore,",
	"In addition,",
	"Meanwhile,",
	"On another note,",
	"Interestingly,",
	"Remarkably,",
	"Specifically,",
	"Notably,",
	"To clarify,",
	"Currently,",
	"Recently,",
}

var emailClosers = []string{
	"Best regards,",
	"Kind regards,",
	"Sincerely,",
	"Thanks,",
	"Cheers,",
	"Best,",
	"Regards,",
	"Warm regards,",
	"Many thanks,",
	"Talk soon,",
}
