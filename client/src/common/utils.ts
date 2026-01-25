export function getRandomIntBetween(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

// public static showAlert(message: string) {
//   alertMessage.textContent = message;
//   alertContainer.classList.remove("hidden");
// }

export function createIndexArray(size: number): number[] {
  return new Array(size).fill(null).map((_: null, index: number) => index);
}

export function shuffleArray(array: unknown[]) {
  return array.sort(() => Math.random() - 0.5);
}

export function isOnline(): boolean {
  return window.navigator.onLine;
}

export function isOccupied(target: HTMLElement) {
  return target.innerText !== "";
}

//   public static generateUUID(): string {
//     return crypto.randomUUID();
//   }
//
