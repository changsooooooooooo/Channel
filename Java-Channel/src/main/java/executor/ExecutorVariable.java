package executor;

import java.util.concurrent.Executors;
import java.util.concurrent.ThreadFactory;

public class ExecutorVariable {
    public final java.util.concurrent.Executor excutable;

    public ExecutorVariable(){
        this.excutable =
                Executors.newCachedThreadPool(
                        new ThreadFactory() {
                            @Override
                            public Thread newThread(Runnable r) {
                                Thread thread = new Thread(r);
                                thread.setDaemon(true);
                                return thread;
                            }
                        }
                );
    }
}
