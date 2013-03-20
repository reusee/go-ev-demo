extern void call_func(void*);

int time_cb(aeEventLoop *loop, long long id, void* func) {
  call_func(func);
  return AE_NOMORE;
}

void create_time_event(aeEventLoop *loop, long long milliseconds, void* func) {
  aeCreateTimeEvent(loop, milliseconds, time_cb, func, NULL);
}
