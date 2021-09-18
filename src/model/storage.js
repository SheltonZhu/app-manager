let storage = {
  set(key, value) {
    localStorage.setItem(key, JSON.stringify(value));
  },
  get(key) {
    return JSON.parse(localStorage.getItem(key));
  },
  getForIndex(index) {
    return localStorage.key(index);
  },
  getKeys() {
    let items = this.getAll();
    let keys = [];
    for (let index = 0; index < items.length; index++) {
      keys.push(items[index].key);
    }
    return keys;
  },
  getLength() {
    return localStorage.length;
  },
  getSupport() {
    return typeof Storage !== "undefined" ? true : false;
  },
  remove(key) {
    localStorage.removeItem(key);
  },
  removeAll() {
    localStorage.clear();
  },
  getAll() {
    let len = this.getLength();
    let arr = new Array();
    for (let i = 0; i < len; i++) {
      let getKey = this.getForIndex(i);
      let getVal = localStorage.getItem(getKey);
      arr[i] = {
        key: getKey,
        val: getVal,
      };
    }
    return arr;
  },
};
export default storage;
