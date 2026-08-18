// this shouldn't be modified, as it's being use at core logic
export default {
  WIN_POSITIONS: [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
    [1, 4, 7],
    [2, 5, 8],
    [3, 6, 9],
    [1, 5, 9],
    [3, 5, 7],
  ],
  WIN_POSITION_OFFSET: -1,
} as const;
