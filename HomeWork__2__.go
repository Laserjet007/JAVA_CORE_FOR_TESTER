//         1. Напишите метод, на вход которого подается двумерный строковый массив размером 4х4, при подаче массива другого размера необходимо
//         бросить исключение MyArraySizeException.
//         2. Далее должен пройтись по всем элементам массива, преобразовать в int и просуммировать. Если в каком-то элементе массива
//         преобразование не удалось (например, в ячейке лежит символ или текст числа), должно быть брошено исключение MyArrayDataException -
//         с детализацией, в какой именно ячейке лежат неверные данные.
//         3. В методе main () вызвать полученный метод, обработать возможные исключения MySizeArrayException и MyArrayDataException и вывести результат расчета.
/**
*Java_Core. Lasson2.0
*
*@author Lazarev Alexei
*@version 20.12.2021
*/

public  class  Main {

    public  static  void  main ( String [] args ) {

        String [] [] arr =  new  String [] [] {{ " 1 " , " 2 " , " 3 " , " 4 " }, { " 2 " , " 2 " , " 2 " , " 3 " }, { " 1 " , " 2 " , " 2 " , " 2 " },{ " 2 " ,« 2 » , « 2 » , « 2 » }};
        попробуйте {
            попробуйте {
                int результат = метод (обр);
                Система . из . println (результат);
            } catch ( MyArraySizeException e) {
                Система . из . println ( " Размер массива превышен! " );
            }
        }
        catch ( MyArrayDataException e) {
            Система . из . println ( " Неправильное значение массива! " );
            Система . из . println ( " Ошибка в ячейке: "  + e . i +  " x "  + e . j);
        }

    }


    общедоступный  статический  метод int  ( String [] [] arr ) выбрасывает MyArraySizeException , MyArrayDataException { 
        int count =  0 ;
        if ( длина обр . ! =  4 ) {
            throw  new  MyArraySizeException ();
        }
        for ( int i =  0 ; i < arr . length; i ++ ) {
            if (arr [i] . length ! =  4 ) {
                throw  new  MyArraySizeException ();
            }
            for ( int j =  0 ; j < arr [i] . length; j ++ ) {
                попробуйте {
                    счетчик = счетчик +  целое число . parseInt (arr [i] [j]);
                }
                catch ( NumberFormatException e) {
                    выбросить  новое  исключение MyArrayDataException (i, j);
                }
            }

        }
        счетчик возврата ;
    }

}
