import { statuses } from "./store";

export default class Socket {
  constructor(server) {
    this.socket = new WebSocket(server);
    this.statuses;
    this.unsubscribe = statuses.subscribe((value) => (this.statuses = value));

    this.socket.onopen = () => {
      this.socket.send(JSON.stringify({ action: "init" }));
    };

    this.socket.onmessage = (event) => {
      let data = JSON.parse(event.data);
      statuses.update((statusStore) => {
        return {
          ...statusStore,
          [data.service]: {
            status: data.status,
            progress: data.progress,
          },
        };
      });
    };
  }
}
