export type Mood = 'success' | 'warning' | 'error';

export type Log = {
  mood: Mood;
  message: string;
};
