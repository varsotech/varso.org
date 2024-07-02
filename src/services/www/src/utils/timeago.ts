// Epochs
const epochs: [string, number][] = [
  ["year", 31536000],
  ["month", 2592000],
  ["day", 86400],
  ["hour", 3600],
  ["minute", 60],
  ["second", 1],
];

// Get duration
const getDuration = (
  timeAgoInSeconds: number
): { interval: number; epoch: string } => {
  for (let [name, seconds] of epochs) {
    const interval = Math.floor(timeAgoInSeconds / seconds);

    if (interval >= 1) {
      return {
        interval: interval,
        epoch: name,
      };
    }
  }

  return {
    interval: 0,
    epoch: "",
  };
};

// Calculate
export const timeAgo = (date: Date) => {
  const currentDate = new Date();
  const differenceInMillis = currentDate.getTime() - new Date(date).getTime();
  const timeAgoInSeconds = differenceInMillis / 1000;

  const { interval, epoch } = getDuration(timeAgoInSeconds);
  const suffix = interval === 1 ? "" : "s";

  return `${interval} ${epoch}${suffix} ago`;
};
