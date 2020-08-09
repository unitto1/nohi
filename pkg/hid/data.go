// copied from docker-ce project
// original source code:
// https://github.com/docker/docker-ce/blob/master/components/engine/pkg/namesgenerator/names-generator.go
package hid

var left = [...]string{
	"admiring",
	"adoring",
	"affectionate",
	"agitated",
	"amazing",
	"angry",
	"awesome",
	"beautiful",
	"blissful",
	"bold",
	"boring",
	"brave",
	"busy",
	"charming",
	"clever",
	"cool",
	"compassionate",
	"competent",
	"condescending",
	"confident",
	"cranky",
	"crazy",
	"dazzling",
	"determined",
	"distracted",
	"dreamy",
	"eager",
	"ecstatic",
	"elastic",
	"elated",
	"elegant",
	"eloquent",
	"epic",
	"exciting",
	"fervent",
	"festive",
	"flamboyant",
	"focused",
	"friendly",
	"frosty",
	"funny",
	"gallant",
	"gifted",
	"goofy",
	"gracious",
	"great",
	"happy",
	"hardcore",
	"heuristic",
	"hopeful",
	"hungry",
	"infallible",
	"inspiring",
	"interesting",
	"intelligent",
	"jolly",
	"jovial",
	"keen",
	"kind",
	"laughing",
	"loving",
	"lucid",
	"magical",
	"mystifying",
	"modest",
	"musing",
	"naughty",
	"nervous",
	"nice",
	"nifty",
	"nostalgic",
	"objective",
	"optimistic",
	"peaceful",
	"pedantic",
	"pensive",
	"practical",
	"priceless",
	"quirky",
	"quizzical",
	"recursing",
	"relaxed",
	"reverent",
	"romantic",
	"sad",
	"serene",
	"sharp",
	"silly",
	"sleepy",
	"stoic",
	"strange",
	"stupefied",
	"suspicious",
	"sweet",
	"tender",
	"thirsty",
	"trusting",
	"unruffled",
	"upbeat",
	"vibrant",
	"vigilant",
	"vigorous",
	"wizardly",
	"wonderful",
	"xenodochial",
	"youthful",
	"zealous",
	"zen",
}

var right = [...]string{
	// https://en.wikipedia.org/wiki/Mu%E1%B8%A5ammad_ibn_J%C4%81bir_al-%E1%B8%A4arr%C4%81n%C4%AB_al-Batt%C4%81n%C4%AB
	"albattani",

	// https://en.wikipedia.org/wiki/Frances_E._Allen
	"allen",

	// https://en.wikipedia.org/wiki/June_Almeida
	"almeida",

	// https://en.wikipedia.org/wiki/Kathleen_Antonelli
	"antonelli",

	// https://en.wikipedia.org/wiki/Maria_Gaetana_Agnesi
	"agnesi",

	// https://en.wikipedia.org/wiki/Archimedes
	"archimedes",

	// https://en.wikipedia.org/wiki/Maria_Ardinghelli
	"ardinghelli",

	// https://en.wikipedia.org/wiki/Aryabhata
	"aryabhata",

	// https://en.wikipedia.org/wiki/Wanda_Austin
	"austin",

	// https://en.wikipedia.org/wiki/Charles_Babbage.
	"babbage",

	// https://en.wikipedia.org/wiki/Stefan_Banach
	"banach",

	// https://en.wikipedia.org/wiki/The_Adventures_of_Buckaroo_Banzai_Across_the_8th_Dimension
	"banzai",

	// https://en.wikipedia.org/wiki/John_Bardeen
	"bardeen",

	// https://en.wikipedia.org/wiki/Jean_Bartik
	"bartik",

	// https://en.wikipedia.org/wiki/Laura_Bassi
	"bassi",

	// https://en.wikipedia.org/wiki/Hugh_Beaver
	"beaver",

	// https://en.wikipedia.org/wiki/Alexander_Graham_Bell
	"bell",

	// https://en.wikipedia.org/wiki/Karl_Benz
	"benz",

	// https://en.wikipedia.org/wiki/Homi_J._Bhabha
	"bhabha",

	// https://en.wikipedia.org/wiki/Bh%C4%81skara_II#Calculus
	"bhaskara",

	// https://en.wikipedia.org/wiki/Sue_Black_(computer_scientist)
	"black",

	// https://en.wikipedia.org/wiki/Elizabeth_Blackburn
	"blackburn",

	// https://en.wikipedia.org/wiki/Elizabeth_Blackwell
	"blackwell",

	// https://en.wikipedia.org/wiki/Niels_Bohr.
	"bohr",

	// https://en.wikipedia.org/wiki/Kathleen_Booth
	"booth",

	// https://en.wikipedia.org/wiki/Anita_Borg
	"borg",

	// https://en.wikipedia.org/wiki/Satyendra_Nath_Bose
	"bose",

	// https://en.wikipedia.org/wiki/Katie_Bouman
	"bouman",

	// https://en.wikipedia.org/wiki/Evelyn_Boyd_Granville
	"boyd",

	// https://en.wikipedia.org/wiki/Brahmagupta#Zero
	"brahmagupta",

	// https://en.wikipedia.org/wiki/Walter_Houser_Brattain
	"brattain",

	// https://en.wikipedia.org/wiki/Emmett_Brown
	"brown",

	// https://en.wikipedia.org/wiki/Linda_B._Buck
	"buck",

	// https://en.wikipedia.org/wiki/Jocelyn_Bell_Burnell
	"burnell",

	// https://en.wikipedia.org/wiki/Annie_Jump_Cannon
	"cannon",

	// https://en.wikipedia.org/wiki/Rachel_Carson
	"carson",

	// https://en.wikipedia.org/wiki/Mary_Cartwright
	"cartwright",

	// https://en.wikipedia.org/wiki/George_Washington_Carver
	"carver",

	// https://en.wikipedia.org/wiki/Vint_Cerf
	"cerf",

	// https://en.wikipedia.org/wiki/Subrahmanyan_Chandrasekhar
	"chandrasekhar",

	// https://en.wikipedia.org/wiki/Sergey_Chaplygin
	"chaplygin",

	// https://en.wikipedia.org/wiki/%C3%89milie_du_Ch%C3%A2telet
	"chatelet",

	// https://en.wikipedia.org/wiki/Asima_Chatterjee
	"chatterjee",

	// https://en.wikipedia.org/wiki/Pafnuty_Chebyshev
	"chebyshev",

	// https://en.wikipedia.org/wiki/Bram_Cohen
	"cohen",

	// https://en.wikipedia.org/wiki/David_Chaum
	"chaum",

	// https://en.wikipedia.org/wiki/Joan_Clarke
	"clarke",

	// https://en.wikipedia.org/wiki/Jane_Colden
	"colden",

	// https://en.wikipedia.org/wiki/Gerty_Cori
	"cori",

	// https://en.wikipedia.org/wiki/Seymour_Cray
	"cray",

	// https://en.wikipedia.org/wiki/Joan_Curran
	// https://en.wikipedia.org/wiki/Samuel_Curran
	"curran",

	// https://en.wikipedia.org/wiki/Marie_Curie.
	"curie",

	// https://en.wikipedia.org/wiki/Charles_Darwin.
	"darwin",

	// https://en.wikipedia.org/wiki/Leonardo_da_Vinci.
	"davinci",

	// https://en.wikipedia.org/wiki/Alexander_Dewdney
	"dewdney",

	// https://en.wikipedia.org/wiki/Satish_Dhawan
	"dhawan",

	// https://en.wikipedia.org/wiki/Whitfield_Diffie
	"diffie",

	// https://en.wikipedia.org/wiki/Edsger_W._Dijkstra.
	"dijkstra",

	// https://en.wikipedia.org/wiki/Paul_Dirac
	"dirac",

	// https://en.wikipedia.org/wiki/Agnes_Meyer_Driscoll
	"driscoll",

	// https://en.wikipedia.org/wiki/Donna_Dubinsky
	"dubinsky",

	// https://en.wikipedia.org/wiki/Annie_Easley
	"easley",

	// https://en.wikipedia.org/wiki/Thomas_Edison
	"edison",

	// https://en.wikipedia.org/wiki/Albert_Einstein
	"einstein",

	// https://en.wikipedia.org/wiki/Alexandra_Elbakyan
	"elbakyan",

	// https://en.wikipedia.org/wiki/Taher_Elgamal
	"elgamal",

	// https://en.wikipedia.org/wiki/Gertrude_Elion
	"elion",

	// https://en.wikipedia.org/wiki/James_H._Ellis
	"ellis",

	// https://en.wikipedia.org/wiki/Douglas_Engelbart
	"engelbart",

	// https://en.wikipedia.org/wiki/Euclid
	"euclid",

	// https://de.wikipedia.org/wiki/Leonhard_Euler
	"euler",

	// https://en.wikipedia.org/wiki/Michael_Faraday
	"faraday",

	// https://en.wikipedia.org/wiki/Horst_Feistel
	"feistel",

	// https://en.wikipedia.org/wiki/Pierre_de_Fermat
	"fermat",

	// https://en.wikipedia.org/wiki/Enrico_Fermi.
	"fermi",

	// https://en.wikipedia.org/wiki/Richard_Feynman
	"feynman",

	// https://en.wikipedia.org/wiki/Benjamin_Franklin
	"franklin",

	// https://en.wikipedia.org/wiki/Yuri_Gagarin
	"gagarin",

	// https://en.wikipedia.org/wiki/Galileo_Galilei
	"galileo",

	// https://en.wikipedia.org/wiki/%C3%89variste_Galois
	"galois",

	// https://en.wikipedia.org/wiki/Kadambini_Ganguly
	"ganguly",

	// https://en.wikipedia.org/wiki/Bill_Gates
	"gates",

	// https://en.wikipedia.org/wiki/Carl_Friedrich_Gauss
	"gauss",

	// https://en.wikipedia.org/wiki/Sophie_Germain
	"germain",

	// https://en.wikipedia.org/wiki/Adele_Goldberg_(computer_scientist)
	"goldberg",

	// https://en.wikipedia.org/wiki/Adele_Goldstine
	"goldstine",

	// https://en.wikipedia.org/wiki/Shafi_Goldwasser
	"goldwasser",

	// James Golick, all around gangster.
	"golick",

	// https://en.wikipedia.org/wiki/Jane_Goodall
	"goodall",

	// https://en.wikipedia.org/wiki/Stephen_Jay_Gould
	"gould",

	// https://en.wikipedia.org/wiki/Carol_W._Greider
	"greider",

	// https://en.wikipedia.org/wiki/Alexander_Grothendieck
	"grothendieck",

	// https://en.wikipedia.org/wiki/Lois_Haibt
	"haibt",

	// https://en.wikipedia.org/wiki/Margaret_Hamilton_(scientist)
	"hamilton",

	// https://en.wikipedia.org/wiki/Caroline_Haslett
	"haslett",

	// https://en.wikipedia.org/wiki/Stephen_Hawking
	"hawking",

	// https://en.wikipedia.org/wiki/Martin_Hellman
	"hellman",

	// https://en.wikipedia.org/wiki/Werner_Heisenberg
	"heisenberg",

	// https://en.wikipedia.org/wiki/Grete_Hermann
	"hermann",

	// https://en.wikipedia.org/wiki/Caroline_Herschel
	"herschel",

	// https://en.wikipedia.org/wiki/Heinrich_Hertz
	"hertz",

	// https://en.wikipedia.org/wiki/Jaroslav_Heyrovsk%C3%BD
	"heyrovsky",

	// https://en.wikipedia.org/wiki/Dorothy_Hodgkin
	"hodgkin",

	// https://en.wikipedia.org/wiki/Douglas_Hofstadter
	"hofstadter",

	// https://en.wikipedia.org/wiki/Erna_Schneider_Hoover
	"hoover",

	// https://en.wikipedia.org/wiki/Grace_Hopper
	"hopper",

	// https://en.wikipedia.org/wiki/Frances_Hugle
	"hugle",

	// https://en.wikipedia.org/wiki/Hypatia
	"hypatia",

	// https://en.wikipedia.org/wiki/Teruko_Ishizaka
	"ishizaka",

	// https://en.wikipedia.org/wiki/Mary_Jackson_(engineer)
	"jackson",

	// https://en.wikipedia.org/wiki/Jang_Yeong-sil
	"jang",

	// https://en.wikipedia.org/wiki/Mae_Jemison
	"jemison",

	// https://en.wikipedia.org/wiki/Jean_Bartik
	"jennings",

	// https://en.wikipedia.org/wiki/Mary_Lou_Jepsen
	"jepsen",

	// https://en.wikipedia.org/wiki/Katherine_Johnson
	"johnson",

	// https://en.wikipedia.org/wiki/Ir%C3%A8ne_Joliot-Curie
	"joliot",

	// https://en.wikipedia.org/wiki/Karen_Sp%C3%A4rck_Jones
	"jones",

	// https://en.wikipedia.org/wiki/A._P._J._Abdul_Kalam
	"kalam",

	// https://en.wikipedia.org/wiki/Sergey_Kapitsa
	"kapitsa",

	// https://en.wikipedia.org/wiki/Susan_Kare
	"kare",

	// https://en.wikipedia.org/wiki/Mstislav_Keldysh
	"keldysh",

	// https://en.wikipedia.org/wiki/Mary_Kenneth_Keller
	"keller",

	// https://en.wikipedia.org/wiki/Johannes_Kepler
	"kepler",

	// https://en.wikipedia.org/wiki/Omar_Khayyam
	"khayyam",

	// https://en.wikipedia.org/wiki/Har_Gobind_Khorana
	"khorana",

	// https://en.wikipedia.org/wiki/Jack_Kilby
	"kilby",

	// https://en.wikipedia.org/wiki/Maria_Margarethe_Kirch
	"kirch",

	// https://en.wikipedia.org/wiki/Donald_Knuth
	"knuth",

	// https://en.wikipedia.org/wiki/Sofia_Kovalevskaya
	"kowalevski",

	// https://en.wikipedia.org/wiki/Marie-Jeanne_de_Lalande
	"lalande",

	// https://en.wikipedia.org/wiki/Hedy_Lamarr
	"lamarr",

	// https://en.wikipedia.org/wiki/Leslie_Lamport
	"lamport",

	// https://en.wikipedia.org/wiki/Mary_Leakey
	"leakey",

	// https://en.wikipedia.org/wiki/Henrietta_Swan_Leavitt
	"leavitt",

	// https://en.wikipedia.org/wiki/Esther_Lederberg
	"lederberg",

	// https://en.wikipedia.org/wiki/Inge_Lehmann
	"lehmann",

	// https://en.wikipedia.org/wiki/Daniel_Lewin
	"lewin",

	// https://en.wikipedia.org/wiki/Ruth_Teitelbaum
	"lichterman",

	// https://en.wikipedia.org/wiki/Barbara_Liskov
	"liskov",

	// https://en.wikipedia.org/wiki/Ada_Lovelace (thanks James Turnbull)
	"lovelace",

	// https://en.wikipedia.org/wiki/Auguste_and_Louis_Lumi%C3%A8re
	"lumiere",

	// https://en.wikipedia.org/wiki/Mah%C4%81v%C4%ABra_(mathematician)
	"mahavira",

	// https://en.wikipedia.org/wiki/Lynn_Margulis
	"margulis",

	// https://en.wikipedia.org/wiki/Yukihiro_Matsumoto
	"matsumoto",

	// https://en.wikipedia.org/wiki/James_Clerk_Maxwell
	"maxwell",

	// https://en.wikipedia.org/wiki/Maria_Mayer
	"mayer",

	// https://en.wikipedia.org/wiki/John_McCarthy_(computer_scientist)
	"mccarthy",

	// https://en.wikipedia.org/wiki/Barbara_McClintock
	"mcclintock",

	// https://en.wikipedia.org/wiki/Anne_McLaren
	"mclaren",

	// https://en.wikipedia.org/wiki/Malcom_McLean
	"mclean",

	// https://en.wikipedia.org/wiki/Kathleen_Antonelli
	"mcnulty",

	// https://en.wikipedia.org/wiki/Gregor_Mendel
	"mendel",

	// https://en.wikipedia.org/wiki/Dmitri_Mendeleev
	"mendeleev",

	// https://en.wikipedia.org/wiki/Lise_Meitner
	"meitner",

	// https://en.wikipedia.org/wiki/Carla_Meninsky
	"meninsky",

	// https://en.wikipedia.org/wiki/Ralph_Merkle
	"merkle",

	// https://en.wikipedia.org/wiki/Johanna_Mestorf
	"mestorf",

	// https://en.wikipedia.org/wiki/Maryam_Mirzakhani
	"mirzakhani",

	// https://en.wikipedia.org/wiki/Gordon_Moore
	"moore",

	// https://en.wikipedia.org/wiki/Samuel_Morse
	"morse",

	// https://en.wikipedia.org/wiki/Ian_Murdock
	"murdock",

	// https://en.wikipedia.org/wiki/May-Britt_Moser
	"moser",

	// https://en.wikipedia.org/wiki/John_Napier
	"napier",

	// https://en.wikipedia.org/wiki/John_Forbes_Nash_Jr.
	"nash",

	// https://en.wikipedia.org/wiki/Von_Neumann_architecture
	"neumann",

	// https://en.wikipedia.org/wiki/Isaac_Newton
	"newton",

	// https://en.wikipedia.org/wiki/Florence_Nightingale#Statistics_and_sanitary_reform
	"nightingale",

	// https://en.wikipedia.org/wiki/Alfred_Nobel
	"nobel",

	// https://en.wikipedia.org/wiki/Emmy_Noether
	"noether",

	// http://www.businessinsider.com/poppy-northcutt-helped-apollo-astronauts-2014-12?op=1
	"northcutt",

	// https://en.wikipedia.org/wiki/Robert_Noyce
	"noyce",

	// https://en.wikipedia.org/wiki/P%C4%81%E1%B9%87ini#Comparison_with_modern_formal_systems
	"panini",

	// https://en.wikipedia.org/wiki/Ambroise_Par%C3%A9
	"pare",

	// https://en.wikipedia.org/wiki/Blaise_Pascal
	"pascal",

	// https://en.wikipedia.org/wiki/Louis_Pasteur.
	"pasteur",

	// https://en.wikipedia.org/wiki/Cecilia_Payne-Gaposchkin
	"payne",

	// https://en.wikipedia.org/wiki/Radia_Perlman
	"perlman",

	// https://en.wikipedia.org/wiki/Rob_Pike
	"pike",

	// https://en.wikipedia.org/wiki/Henri_Poincar%C3%A9
	"poincare",

	// https://en.wikipedia.org/wiki/Laura_Poitras
	"poitras",

	// https://en.wikipedia.org/wiki/Tatiana_Proskouriakoff
	"proskuriakova",

	// https://en.wikipedia.org/wiki/Ptolemy
	"ptolemy",

	// https://en.wikipedia.org/wiki/C._V._Raman
	"raman",

	// https://en.wikipedia.org/wiki/Srinivasa_Ramanujan
	"ramanujan",

	// https://en.wikipedia.org/wiki/Sally_Ride
	"ride",

	// https://en.wikipedia.org/wiki/Rita_Levi-Montalcini)
	"montalcini",

	// https://en.wikipedia.org/wiki/Dennis_Ritchie
	"ritchie",

	// https://en.wikipedia.org/wiki/Ida_Rhodes
	"rhodes",

	// https://en.wikipedia.org/wiki/Julia_Robinson
	"robinson",

	// https://en.wikipedia.org/wiki/Wilhelm_R%C3%B6ntgen
	"roentgen",

	// https://en.wikipedia.org/wiki/Rosalind_Franklin
	"rosalind",

	// https://en.wikipedia.org/wiki/Vera_Rubin
	"rubin",

	// https://en.wikipedia.org/wiki/Meghnad_Saha
	"saha",

	// https://en.wikipedia.org/wiki/Jean_E._Sammet
	"sammet",

	// https://en.wikipedia.org/wiki/Mildred_Sanderson
	"sanderson",

	// https://en.wikipedia.org/wiki/Satoshi_Nakamoto
	"satoshi",

	// https://en.wikipedia.org/wiki/Adi_Shamir
	"shamir",

	// https://en.wikipedia.org/wiki/Claude_Shannon)
	"shannon",

	// https://en.wikipedia.org/wiki/Carol_Shaw_(video_game_designer)
	"shaw",

	// https://en.wikipedia.org/wiki/Steve_Shirley
	"shirley",

	// https://en.wikipedia.org/wiki/William_Shockley
	"shockley",

	// https://en.wikipedia.org/wiki/Lina_Stern
	"shtern",

	// https://en.wikipedia.org/wiki/Fran%C3%A7oise_Barr%C3%A9-Sinoussi
	"sinoussi",

	// https://en.wikipedia.org/wiki/Betty_Holberton
	"snyder",

	// https://en.wikipedia.org/wiki/Cynthia_Solomon
	"solomon",

	// https://en.wikipedia.org/wiki/Frances_Spence
	"spence",

	// https://en.wikipedia.org/wiki/Michael_Stonebraker
	"stonebraker",

	// https://en.wikipedia.org/wiki/Ivan_Sutherland
	"sutherland",

	// https://en.wikipedia.org/wiki/Janese_Swanson
	"swanson",

	// https://en.wikiquote.org/wiki/Aaron_Swartz
	"swartz",

	// https://en.wikipedia.org/wiki/Bertha_Swirles
	"swirles",

	// https://en.wikipedia.org/wiki/Helen_B._Taussig
	"taussig",

	// https://en.wikipedia.org/wiki/Valentina_Tereshkova
	"tereshkova",

	// https://en.wikipedia.org/wiki/Nikola_Tesla
	"tesla",

	// https://en.wikipedia.org/wiki/Marie_Tharp
	"tharp",

	// https://en.wikipedia.org/wiki/Ken_Thompson
	"thompson",

	// https://en.wikipedia.org/wiki/Linus_Torvalds
	"torvalds",

	// https://en.wikipedia.org/wiki/Tu_Youyou
	"tu",

	// https://en.wikipedia.org/wiki/Alan_Turing.
	"turing",

	// https://en.wikipedia.org/wiki/Var%C4%81hamihira#Contributions
	"varahamihira",

	// https://en.wikipedia.org/wiki/Dorothy_Vaughan
	"vaughan",

	// https://en.wikipedia.org/wiki/Visvesvaraya
	"visvesvaraya",

	// https://en.wikipedia.org/wiki/Christiane_N%C3%BCsslein-Volhard
	"volhard",

	// https://en.wikipedia.org/wiki/C%C3%A9dric_Villani
	"villani",

	// https://en.wikipedia.org/wiki/Marlyn_Meltzer
	"wescoff",

	// https://en.wikipedia.org/wiki/Sylvia_Wilbur
	"wilbur",

	// https://en.wikipedia.org/wiki/Andrew_Wiles
	"wiles",

	// https://en.wikipedia.org/wiki/Roberta_Williams
	"williams",

	// https://en.wikipedia.org/wiki/Malcolm_J._Williamson
	"williamson",

	// https://en.wikipedia.org/wiki/Sophie_Wilson
	"wilson",

	// https://en.wikipedia.org/wiki/Jeannette_Wing
	"wing",

	// https://en.wikipedia.org/wiki/Steve_Wozniak
	"wozniak",

	// https://en.wikipedia.org/wiki/Wright_brothers
	"wright",

	// https://en.wikipedia.org/wiki/Chien-Shiung_Wu
	"wu",

	// https://en.wikipedia.org/wiki/Rosalyn_Sussman_Yalow
	"yalow",

	// https://en.wikipedia.org/wiki/Ada_Yonath
	"yonath",

	// https://en.wikipedia.org/wiki/Nikolay_Yegorovich_Zhukovsky
	"zhukovsky",
}
