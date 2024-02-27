# Calculo de Estabilidade de Equações

## Propósito:
#### O desenvolvimento de uma API, que verifique a estabilidade de um sistema discreto, por meio do Critério de Jury.

## Regras de Negócio
##### considere: "i0k0l1" o primeiro elemento (i0), da segunda matriz (l1), do primeiro passo (k0); "inknln" qualquer elemento (in), de qualquer matriz (ln), de qualquer passo (kn); "jkn" o coeficiente "J" de qualquer passo (kn).
- A API deve receber como entrada os indices equação que descreve o sistema;
- O calculo para determinar a estabilidade da equação será realizado em passos, estes  serão indicados pela variavel "K", onde "K=0" é o primeiro passo e "K=n" é o ultimo passo;
- A quantidade de passos será determinada pelo expoente que indica o grau da equação supracitada, acrescido de uma unidade;
- Cada passo será composto de no máximo duas matrizes linhas e um coeficiente "J";
- O coeficiente "J" é calculado por meio do modulo da divisão i0knl2/i0knl1;
- A primeira matriz do primeiro passo (K=0), será a seguencia dos indices da equação, determinada pela ordem decrescente dos expoentes, e a segunda será a inversão da seguencia da matriz anterior (a segunda matriz seguirá esta regra em todos os passos);
- A partir do segundo passo (K=1), a primeira matriz será calculada da seguinte forma: inknl1 = ink(n-1)l1 - ink(n-1)l2 * jkn;
- Caso a equação que descreve o sistema seja um polinômio incompleto, os indices faltantes devem ser substituidos por zero;
- A partir do segundo passo (k=1), o calculo irá zerar o ultimo elemento da primeira matriz, visto isso, na criação da segunda matriz, este elemento deve ser desconsiderado, utilizando-se do imediatamente anterior;
- Ao finalizar o calculo é esperado que o ultimo passo (k=n), possua apenas uma matriz (1x1);
- Condições para instabilidade do sistema: jkn >= 1 ; o único elemento da matriz do ultimo passo é igual a zero; se ocorrer divisão por 0 em qualquer passo;
- Como reposta ao usuário a api deve retornar: a informação se a equação é ou não estavel e os resultados de cada passo do calculo.