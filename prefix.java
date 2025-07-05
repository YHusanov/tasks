public class Main {
    public static String longestCommonPrefix(String[] strs) {
        if (strs.length == 0) {
            return "";
        }

        String prefix = strs[0];

        for (int i = 1; i < strs.length; i++) {
            while (prefix.length() > 0 &&
                    (strs[i].length() < prefix.length() ||
                            !strs[i].substring(0, prefix.length()).equals(prefix))) {
                prefix = prefix.substring(0, prefix.length() - 1);
            }

            if (prefix.isEmpty()) {
                return "";
            }
        }

        return prefix;
    }

    public static void main(String[] args) {
        String[] strs = {"flower", "flow", "flight"};

        String result = longestCommonPrefix(strs);

        if (result.isEmpty()) {
            System.out.println("Prefix mavjud emas.");
        } else {
            System.out.println("Eng uzun umumiy prefix: " + result);
        }
    }
}
