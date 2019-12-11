#include <cstdio>
#include <cstdlib>
#include <vector>

using vec = std::vector<int>;

struct cpu {
	vec M = { };
	cpu(FILE *fp) {
		for (int k; fscanf(fp, "%d,", &k) == 1; M.push_back(k));
	}
	auto run(int phase) {
		vec A, M = this->M;
		for (int ip = 0; ; ) switch (M[ip]) {
		    default: printf("unimplemented opcode %d\n", M[ip]); abort();
		    case   99: return A;
		    case  105: ip = M[ip+1] ? M[M[ip+2]] : ip + 3;            break;
		    case  101: M[M[ip+3]] =   M[ip+1]  + M[M[ip+2]]; ip += 4; break;
		    case 1001: M[M[ip+3]] = M[M[ip+1]] +   M[ip+2] ; ip += 4; break;
		    case  102: M[M[ip+3]] =   M[ip+1]  * M[M[ip+2]]; ip += 4; break;
		    case 1002: M[M[ip+3]] = M[M[ip+1]] *   M[ip+2] ; ip += 4; break;
		    case    3: M[M[ip+1]] = phase; phase = 0;        ip += 2; break;
		    case    4: A.push_back(M[M[ip+1]]);              ip += 2; break;
		}
	}
};

int main(int argc, char **argv) {
	if (argc < 2) {
		printf("usage: %s INPUT\n", *argv);
		return 1;
	}

	FILE *fp = fopen(argv[1], "r");
	if (!fp) {
		printf("Error opening file\n");
		return 1;
	}

	cpu C(fp);

	std::vector<vec> Amp;
	vec Weight, MulMask;
	for (int i = 0; i < 5; i++) {
		auto a = C.run(i + 5);
		Amp.push_back(a);

		Weight.resize(a.size(), 1);
		MulMask.resize(a.size());
		for (int j = 0; j < a.size(); j++) {
			MulMask[j] |= !a[j] << i;
			if (j && !a[j]) Weight[j-1] *= 2;
		}
	}

	for (int i = Weight.size() - 1; i >= 1; i--) {
		Weight[i-1] *= Weight[i];
	}

	vec best(32);
	for (int m = 1; m < 32; m++) {

		for (int bi = 0, b = 1; bi < 5; bi++, b *= 2) {

			if (~m & b) continue;

			int value = best[m ^ b];

			for (int i = 0; i < Amp[bi].size(); i++) {
				int s = __builtin_popcount(MulMask[i] & m);
				value += Amp[bi][i] * (Weight[i] << s);
			}
			best[m] = std::max(best[m], value);
		}
	}

	int part2 = best.back();
	printf("Part 2: %d\n", part2);

	return 0;
}