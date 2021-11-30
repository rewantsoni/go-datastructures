package list

import (
	"errors"
	"github.com/rewantsoni/go-datastructures/iterator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewArrayList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() *ArrayList
		expectedResult *ArrayList
		expectedError  error
	}{
		{
			name: "test create new empty array list",
			actualResult: func() *ArrayList {
				return NewArrayList()
			},
			expectedResult: &ArrayList{
				size:            0,
				capacity:        16,
				scalingFactor:   2,
				upperLoadFactor: 0.75,
				lowerLoadFactor: 0.4,
				data:            make([]int, 16),
			},
		},
		{
			name: "test create new array list with elements",
			actualResult: func() *ArrayList {
				return NewArrayList(1, 2, 3, 4, 5)
			},
			expectedResult: &ArrayList{
				size:            5,
				capacity:        16,
				scalingFactor:   2,
				upperLoadFactor: 0.75,
				lowerLoadFactor: 0.4,
				data:            []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			name: "test create new array list with 1000 elements",
			actualResult: func() *ArrayList {
				data := make([]int, 1000)
				for i := 0; i < 1000; i++ {
					data[i] = i
				}

				return NewArrayList(data...)
			},
			expectedResult: &ArrayList{
				size:            1000,
				capacity:        2048,
				scalingFactor:   2,
				upperLoadFactor: 0.75,
				lowerLoadFactor: 0.4,
				data:            []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 384, 385, 386, 387, 388, 389, 390, 391, 392, 393, 394, 395, 396, 397, 398, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 446, 447, 448, 449, 450, 451, 452, 453, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500, 501, 502, 503, 504, 505, 506, 507, 508, 509, 510, 511, 512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 524, 525, 526, 527, 528, 529, 530, 531, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 550, 551, 552, 553, 554, 555, 556, 557, 558, 559, 560, 561, 562, 563, 564, 565, 566, 567, 568, 569, 570, 571, 572, 573, 574, 575, 576, 577, 578, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 590, 591, 592, 593, 594, 595, 596, 597, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 608, 609, 610, 611, 612, 613, 614, 615, 616, 617, 618, 619, 620, 621, 622, 623, 624, 625, 626, 627, 628, 629, 630, 631, 632, 633, 634, 635, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 658, 659, 660, 661, 662, 663, 664, 665, 666, 667, 668, 669, 670, 671, 672, 673, 674, 675, 676, 677, 678, 679, 680, 681, 682, 683, 684, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 695, 696, 697, 698, 699, 700, 701, 702, 703, 704, 705, 706, 707, 708, 709, 710, 711, 712, 713, 714, 715, 716, 717, 718, 719, 720, 721, 722, 723, 724, 725, 726, 727, 728, 729, 730, 731, 732, 733, 734, 735, 736, 737, 738, 739, 740, 741, 742, 743, 744, 745, 746, 747, 748, 749, 750, 751, 752, 753, 754, 755, 756, 757, 758, 759, 760, 761, 762, 763, 764, 765, 766, 767, 768, 769, 770, 771, 772, 773, 774, 775, 776, 777, 778, 779, 780, 781, 782, 783, 784, 785, 786, 787, 788, 789, 790, 791, 792, 793, 794, 795, 796, 797, 798, 799, 800, 801, 802, 803, 804, 805, 806, 807, 808, 809, 810, 811, 812, 813, 814, 815, 816, 817, 818, 819, 820, 821, 822, 823, 824, 825, 826, 827, 828, 829, 830, 831, 832, 833, 834, 835, 836, 837, 838, 839, 840, 841, 842, 843, 844, 845, 846, 847, 848, 849, 850, 851, 852, 853, 854, 855, 856, 857, 858, 859, 860, 861, 862, 863, 864, 865, 866, 867, 868, 869, 870, 871, 872, 873, 874, 875, 876, 877, 878, 879, 880, 881, 882, 883, 884, 885, 886, 887, 888, 889, 890, 891, 892, 893, 894, 895, 896, 897, 898, 899, 900, 901, 902, 903, 904, 905, 906, 907, 908, 909, 910, 911, 912, 913, 914, 915, 916, 917, 918, 919, 920, 921, 922, 923, 924, 925, 926, 927, 928, 929, 930, 931, 932, 933, 934, 935, 936, 937, 938, 939, 940, 941, 942, 943, 944, 945, 946, 947, 948, 949, 950, 951, 952, 953, 954, 955, 956, 957, 958, 959, 960, 961, 962, 963, 964, 965, 966, 967, 968, 969, 970, 971, 972, 973, 974, 975, 976, 977, 978, 979, 980, 981, 982, 983, 984, 985, 986, 987, 988, 989, 990, 991, 992, 993, 994, 995, 996, 997, 998, 999, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test size on creating an new list",
			actualResult: func() int {
				al := NewArrayList()
				return al.size
			},
			expectedResult: 0,
		},
		{
			name: "test size on creating an new list with elements",
			actualResult: func() int {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.size
			},
			expectedResult: 5,
		},
		{
			name: "test size on creating an new list with 100 elements",
			actualResult: func() int {
				data := make([]int, 100)
				for i := 0; i < 100; i++ {
					data[i] = i
				}
				al := NewArrayList(data...)
				return al.size
			},
			expectedResult: 100,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			size := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, size)
		})
	}
}

func TestArrayListAdd(t *testing.T) {
	testCases := []struct {
		name              string
		actualResult      func() (int, bool, *ArrayList)
		expectedArrayList func() *ArrayList
		expectedResult    bool
		expectedSize      int
	}{
		{
			name: "test Size is 1 after adding one element",
			actualResult: func() (int, bool, *ArrayList) {
				al := NewArrayList()
				res := al.Add(1)
				return al.Size(), res, al
			},
			expectedSize: 1,
			expectedArrayList: func() *ArrayList {
				al := &ArrayList{
					size:            1,
					capacity:        16,
					scalingFactor:   2,
					upperLoadFactor: 0.75,
					lowerLoadFactor: 0.4,
					data:            []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}
				return al
			},
			expectedResult: true,
		},
		{
			name: "test Size is 2 after adding two element",
			actualResult: func() (int, bool, *ArrayList) {
				al := NewArrayList()

				res := al.Add(1)

				res = al.Add(2)
				return al.Size(), res, al
			},
			expectedSize: 2,
			expectedArrayList: func() *ArrayList {
				al := &ArrayList{
					size:            2,
					capacity:        16,
					scalingFactor:   2,
					upperLoadFactor: 0.75,
					lowerLoadFactor: 0.4,
					data:            []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}
				return al
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			size, res, resArray := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedSize, size)
			assert.Equal(t, testCase.expectedArrayList(), resArray)
		})
	}
}

func TestArrayListAddAll(t *testing.T) {
	testCases := []struct {
		name              string
		actualResult      func() (int, bool, *ArrayList)
		expectedArrayList func() *ArrayList
		expectedResult    bool
		expectedSize      int
	}{
		{
			name: "test Size is 5 after adding five element",
			actualResult: func() (int, bool, *ArrayList) {
				al := NewArrayList()
				res := al.AddAll(1, 2, 3, 4, 5)
				return al.Size(), res, al
			},
			expectedSize: 5,
			expectedArrayList: func() *ArrayList {
				al := &ArrayList{
					size:            5,
					capacity:        16,
					scalingFactor:   2,
					upperLoadFactor: 0.75,
					lowerLoadFactor: 0.4,
					data:            []int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}
				return al
			},
			expectedResult: true,
		},
		{
			name: "test Size is 17 after adding seventeen element",
			actualResult: func() (int, bool, *ArrayList) {
				al := NewArrayList()
				data := make([]int, 17)
				for i := 0; i < 17; i++ {
					data[i] = i
				}
				res := al.AddAll(data...)
				return al.Size(), res, al
			},
			expectedSize: 17,
			expectedArrayList: func() *ArrayList {
				al := &ArrayList{
					size:            17,
					capacity:        32,
					scalingFactor:   2,
					upperLoadFactor: 0.75,
					lowerLoadFactor: 0.4,
					data:            []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}
				return al
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			size, res, resArrayList := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedSize, size)
			assert.Equal(t, testCase.expectedArrayList(), resArrayList)
		})
	}
}

func TestArrayListAddAt(t *testing.T) {
	testCases := []struct {
		name              string
		actualResult      func() (bool, *ArrayList)
		expectedArrayList func() *ArrayList
		expectedResult    bool
	}{
		{
			name: "test addAt index 10 when list is empty",
			actualResult: func() (bool, *ArrayList) {
				al := NewArrayList()
				res := al.AddAt(10, 1)
				return res, al
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList()
			},
			expectedResult: false,
		},
		{
			name: "test addAt index 0 when list is empty",
			actualResult: func() (bool, *ArrayList) {
				al := NewArrayList()
				res := al.AddAt(0, 1)
				return res, al
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList(1)
			},
			expectedResult: true,
		},
		{
			name: "test addAt is first index when list is not empty",
			actualResult: func() (bool, *ArrayList) {
				al := NewArrayList(1, 2, 3)

				res := al.AddAt(0, 10)
				return res, al
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList(10, 1, 2, 3)
			},
			expectedResult: true,
		},
		{
			name: "test addAt is last index when list is not empty",
			actualResult: func() (bool, *ArrayList) {
				al := NewArrayList(1, 2, 3)

				res := al.AddAt(3, 10)
				return res, al
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList(1, 2, 3, 10)
			},
			expectedResult: true,
		},
		{
			name: "test addAt is last index+1 when list is not empty",
			actualResult: func() (bool, *ArrayList) {
				al := NewArrayList(1, 2, 3)

				res := al.AddAt(4, 10)
				return res, al
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList(1, 2, 3)
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, resArrayList := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedArrayList(), resArrayList)
		})
	}
}

func TestArrayListGetAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int, bool)
		expectedResult int
		expectedBool   bool
	}{
		{
			name: "test get when list is empty",
			actualResult: func() (int, bool) {
				al := NewArrayList()
				return al.GetAt(0)
			},
			expectedResult: -1,
			expectedBool:   false,
		},
		{
			name: "test get first index when list has elements",
			actualResult: func() (int, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.GetAt(0)
			},
			expectedResult: 1,
			expectedBool:   true,
		},
		{
			name: "test get last index when list has elements",
			actualResult: func() (int, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.GetAt(4)
			},
			expectedResult: 5,
			expectedBool:   true,
		},
		{
			name: "test get when index is out of bond and positive",
			actualResult: func() (int, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.GetAt(5)
			},
			expectedResult: -1,
			expectedBool:   false,
		},
		{
			name: "test get when index is out of bond and negative",
			actualResult: func() (int, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.GetAt(-1)
			},
			expectedResult: -1,
			expectedBool:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, b := testCase.actualResult()

			assert.Equal(t, testCase.expectedBool, b)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListContains(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test when list is empty",
			actualResult: func() bool {
				al := NewArrayList()
				return al.Contains(1)
			},
			expectedResult: false,
		},
		{
			name: "test when list contains the element",
			actualResult: func() bool {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.Contains(1)
			},
			expectedResult: true,
		},
		{
			name: "test when list does not contain the element",
			actualResult: func() bool {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.Contains(10)
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListContainsAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test when list is empty",
			actualResult: func() bool {
				al := NewArrayList()
				return al.ContainsAll(1, 2, 3)
			},
			expectedResult: false,
		},
		{
			name: "test when list contains all the elements",
			actualResult: func() bool {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.ContainsAll(1, 2, 5)
			},
			expectedResult: true,
		},
		{
			name: "test when list does not contain the elements",
			actualResult: func() bool {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.ContainsAll(10, 1)
			},
			expectedResult: false,
		},
		{
			name: "test when list does not contain the elements",
			actualResult: func() bool {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.ContainsAll(1, 3, 10)
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListIndexOf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
		expectedError  error
	}{
		{
			name: "test indexOf when list is empty",
			actualResult: func() int {
				al := NewArrayList()
				return al.IndexOf(1)
			},
			expectedResult: -1,
			expectedError:  errors.New("invalid operation: list is empty"),
		},
		{
			name: "test indexOf when element at first index",
			actualResult: func() int {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.IndexOf(1)
			},
			expectedResult: 0,
			expectedError:  nil,
		},
		{
			name: "test indexOf when element at last index",
			actualResult: func() int {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.IndexOf(5)
			},
			expectedResult: 4,
			expectedError:  nil,
		},
		{
			name: "test indexOf when element not in list",
			actualResult: func() int {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.IndexOf(6)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 6 not found in the arrayList"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListReplace(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (*ArrayList, bool)
		expectedResult func() *ArrayList
		expectedBool   bool
	}{
		{
			name: "test replace when arrayList is empty",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList()
				err := al.Replace(1, 4)
				return al, err
			},
			expectedResult: func() *ArrayList {
				al := NewArrayList()
				return al
			},
			expectedBool: false,
		},
		{
			name: "test replace when arrayList has elements",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)
				res := al.Replace(1, 4)
				return al, res
			},
			expectedResult: func() *ArrayList {
				return NewArrayList(4, 2, 3, 4, 5)
			},
			expectedBool: true,
		},
		{
			name: "test replace when arrayList has same element twice",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList(1, 2, 1, 4, 5)
				res := al.Replace(1, 4)
				return al, res
			},
			expectedResult: func() *ArrayList {
				return NewArrayList(4, 2, 1, 4, 5)
			},
			expectedBool: true,
		},
		{
			name: "test replace when arrayList does not have element",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList(2, 2, 3, 4, 5)
				res := al.Replace(1, 4)
				return al, res
			},
			expectedResult: func() *ArrayList {
				return NewArrayList(2, 2, 3, 4, 5)
			},
			expectedBool: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, b := testCase.actualResult()
			assert.Equal(t, testCase.expectedBool, b)
			assert.Equal(t, testCase.expectedResult(), res)
		})

	}
}

func TestArrayListSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (*ArrayList, bool)
		expectedResult func() *ArrayList
		expectedBool   bool
	}{
		{
			name: "test Set when arrayList is empty",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList()
				res := al.Set(1, 4)
				return al, res
			},
			expectedBool: false,
			expectedResult: func() *ArrayList {
				return NewArrayList()
			},
		},
		{
			name: "test Set when arrayList has elements",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)
				res := al.Set(1, 4)
				return al, res
			},
			expectedBool: true,
			expectedResult: func() *ArrayList {
				return NewArrayList(1, 4, 3, 4, 5)
			},
		},
		{
			name: "test Set when index is out of bond",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)
				res := al.Set(10, 4)
				return al, res
			},
			expectedBool: false,
			expectedResult: func() *ArrayList {
				return NewArrayList(1, 2, 3, 4, 5)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, b := testCase.actualResult()
			assert.Equal(t, testCase.expectedBool, b)
			assert.Equal(t, testCase.expectedResult(), res)
		})

	}
}

func TestArrayListRemove(t *testing.T) {
	testCases := []struct {
		name              string
		actualResult      func() (*ArrayList, bool)
		expectedResult    bool
		expectedArrayList func() *ArrayList
	}{
		{
			name: "test remove element when list is empty",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList()

				res := al.Remove(1)
				return al, res
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList()
			},
			expectedResult: false,
		},
		{
			name: "test remove first occurrence of element",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList(1, 2, 3, 1, 5)

				res := al.Remove(1)
				return al, res
			},
			expectedArrayList: func() *ArrayList {
				al := NewArrayList(2, 3, 1, 5)
				return al
			},
			expectedResult: true,
		},
		{
			name: "test remove element when element not present",
			actualResult: func() (*ArrayList, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)
				res := al.Remove(6)
				return al, res
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList(1, 2, 3, 4, 5)
			},
			expectedResult: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			al, res := testCase.actualResult()
			assert.Equal(t, testCase.expectedArrayList(), al)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRemoveAt(t *testing.T) {
	testCases := []struct {
		name              string
		actualResult      func() (*ArrayList, int, bool)
		expectedResult    int
		expectedArrayList func() *ArrayList
		expectedBool      bool
	}{
		{
			name: "test removeAt element when list is empty",
			actualResult: func() (*ArrayList, int, bool) {
				al := NewArrayList()

				res, err := al.RemoveAt(1)
				return al, res, err
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList()
			},
			expectedBool:   false,
			expectedResult: -1,
		},
		{
			name: "test removeAt index when list has elements",
			actualResult: func() (*ArrayList, int, bool) {
				al := NewArrayList(1, 2, 3, 1, 5)

				res, err := al.RemoveAt(1)
				return al, res, err
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList(1, 3, 1, 5)
			},
			expectedBool:   true,
			expectedResult: 2,
		},
		{
			name: "test removeAt element when index not present",
			actualResult: func() (*ArrayList, int, bool) {
				al := NewArrayList(1, 2, 3, 4, 5)

				res, err := al.RemoveAt(6)
				return al, res, err
			},
			expectedArrayList: func() *ArrayList {
				return NewArrayList(1, 2, 3, 4, 5)
			},
			expectedBool:   false,
			expectedResult: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			al, res, b := testCase.actualResult()
			assert.Equal(t, testCase.expectedBool, b)
			assert.Equal(t, testCase.expectedArrayList(), al)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRetainAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() *ArrayList
		expectedResult func() *ArrayList
	}{
		{
			name: "test when arrayList is empty",
			actualResult: func() *ArrayList {
				al := NewArrayList()
				al.RetainAll(1, 2, 3)
				return al
			},
			expectedResult: func() *ArrayList {
				return NewArrayList()
			},
		},
		{
			name: "test when arrayList is not empty",
			actualResult: func() *ArrayList {
				al := NewArrayList(1, 2, 3, 4, 5, 1)
				al.RetainAll(1, 2, 3)
				return al
			},
			expectedResult: func() *ArrayList {
				return NewArrayList(1, 2, 3, 1)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestArrayListRemoveAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() *ArrayList
		expectedResult func() *ArrayList
	}{
		{
			name: "test when arrayList is empty",
			actualResult: func() *ArrayList {
				al := NewArrayList()
				al.RemoveAll(1, 2, 3)
				return al
			},
			expectedResult: func() *ArrayList {
				return NewArrayList()
			},
		},
		{
			name: "test when arrayList is not empty",
			actualResult: func() *ArrayList {
				al := NewArrayList(1, 2, 3, 4, 5, 1)
				al.RemoveAll(1, 2, 3)
				return al
			},
			expectedResult: func() *ArrayList {
				return NewArrayList(4,5)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestArrayListReplaceAll(t *testing.T) {
	testMultiply := func(e int) int {
		return e * 2
	}
	testAdd := func(e int) int {
		return e + 10
	}

	testCases := []struct {
		name           string
		actualResult   func() *ArrayList
		expectedResult func() *ArrayList
	}{
		{
			name: "test replace all when array is empty",
			actualResult: func() *ArrayList {
				al := NewArrayList()
				al.ReplaceAll(testMultiply)
				return al
			},
			expectedResult: func() *ArrayList {
				return NewArrayList()
			},
		},
		{
			name: "test replace all when function multiply and array not empty",
			actualResult: func() *ArrayList {
				al := NewArrayList(1, 2, 3)
				al.ReplaceAll(testMultiply)
				return al
			},
			expectedResult: func() *ArrayList {
				return NewArrayList(2, 4, 6)
			},
		},
		{
			name: "test replace all when add function and array not empty",
			actualResult: func() *ArrayList {
				al := NewArrayList(1, 2, 3)
				al.ReplaceAll(testAdd)
				return al
			},
			expectedResult: func() *ArrayList {
				return NewArrayList(11, 12, 13)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestArrayListIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() iterator.Iterator
		expectedResult iterator.Iterator
	}{
		{
			name: "test Iterator with empty array",
			actualResult: func() iterator.Iterator {
				al := NewArrayList()
				return al.Iterator()
			},
			expectedResult: &arrayListIterator{
				currentIndex: 0,
				al:           NewArrayList(),
			},
		},
		{
			name: "test Iterator with elements",
			actualResult: func() iterator.Iterator {
				al := NewArrayList(1, 2, 3, 4, 5)
				return al.Iterator()
			},
			expectedResult: &arrayListIterator{
				currentIndex: 0,
				al:           NewArrayList(1, 2, 3, 4, 5),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListIteratorHasNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test IteratorHasNext with empty array",
			actualResult: func() bool {
				al := NewArrayList()
				it := al.Iterator()
				return it.HasNext()
			},
			expectedResult: false,
		},
		{
			name: "test Iterator with elements",
			actualResult: func() bool {
				al := NewArrayList(1, 2, 3, 4, 5)
				it := al.Iterator()
				return it.HasNext()
			},

			expectedResult: true,
		},
		{
			name: "test Iterator with elements after array is empty",
			actualResult: func() bool {
				al := NewArrayList(1, 2, 3, 4, 5)
				it := al.Iterator()
				it.Next()
				it.Next()
				it.Next()
				it.Next()
				it.Next()
				return it.HasNext()
			},

			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListIteratorNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test when list is empty",
			actualResult: func() interface{} {
				al := NewArrayList()
				it := al.Iterator()
				var res []interface{}
				res = append(res, it.Next())
				return res
			},
			expectedResult: []interface{}{interface{}(nil)},
		},
		{
			name: "test when list is not empty",
			actualResult: func() interface{} {
				al := NewArrayList(1, 2, 3, 4, 5)
				it := al.Iterator()
				return it.Next()
			},
			expectedResult: 1,
		},
		{
			name: "test when list is not empty get all the element",
			actualResult: func() interface{} {
				al := NewArrayList(1, 2, 3, 4, 5)
				it := al.Iterator()
				var res []interface{}
				res = append(res, it.Next())
				res = append(res, it.Next())
				res = append(res, it.Next())
				res = append(res, it.Next())
				res = append(res, it.Next())
				return res
			},
			expectedResult: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name: "test when list is not empty get extra elements",
			actualResult: func() interface{} {
				al := NewArrayList(1, 2, 3, 4, 5)
				it := al.Iterator()
				var res []interface{}
				res = append(res, it.Next())
				res = append(res, it.Next())
				res = append(res, it.Next())
				res = append(res, it.Next())
				res = append(res, it.Next())
				res = append(res, it.Next())
				return res
			},
			expectedResult: []interface{}{1, 2, 3, 4, 5, interface{}(nil)},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
