package ir.ms.pb;

/**
 * Created by Hamid on 11/29/2017.
 */

public class Log {
	public static ILog logImple;
	public static void d(String module, String debuIfo){
		if(logImple != null){
			logImple.d(module,debuIfo);
		}
	}

	public static interface ILog {
		void d(String module, String debuIfo);
	}
}
