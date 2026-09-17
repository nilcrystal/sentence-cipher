package sentencecipher

// Word lists with 256 words each - Theme: schizo theories / psionics / psychotronics /
// radiations / astral planes / curses / aliens / hyperspace / phantoms / cosmic delirium
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
