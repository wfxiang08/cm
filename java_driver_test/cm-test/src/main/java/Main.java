import com.mysql.jdbc.Connection;

import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

/**
 * Created by dongxu on 1/19/15.
 */
public class Main {
    public static void main(String[] args) throws SQLException, InterruptedException {
        try{
            Class.forName("com.mysql.jdbc.Driver");
        }catch(ClassNotFoundException e1){
            e1.printStackTrace();
        }

        Connection conn = null;
        String url="jdbc:MySQL://127.0.0.1:4000/test?user=root&password=&characterEncoding=utf8&autoReconnect=true";
        try {
            conn = (Connection) DriverManager.getConnection(url);
        } catch (SQLException e){
            e.printStackTrace();
        }

        PreparedStatement ps = conn.prepareStatement("insert into tbl_test values(?, ?)");
        ps.setInt(1, 0);
        ps.setString(2, "hello");
        ps.execute();

        ps = conn.prepareStatement("select * from tbl_test where id = 0");
        ResultSet rs = ps.executeQuery();
        while(rs.next()){
            int id = rs.getInt("id");
            String val = rs.getString("data");

            System.out.println(id);
            System.out.println(val);

            if (id != 0 || val.equals("hello") == false) {
                System.err.println("insert value error");
                System.exit(-1);
            }
        }

        ps = conn.prepareStatement("update tbl_test set data = 'bye' where id = 0");
        int affected = ps.executeUpdate();
        if (affected == 0) {
            System.err.println("update error");
            System.exit(-1);
        }

        ps = conn.prepareStatement("delete from tbl_test where id = 0");
        affected = ps.executeUpdate();
        if  (affected == 0) {
            System.err.println("delete error");
            System.exit(-1);
        }
    }
}
