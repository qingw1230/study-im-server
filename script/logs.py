import os
import shutil
from datetime import datetime

# 设置日志文件所在的目录
log_directory = '../logs'
target_root = '../logs'

# 获取今天的日期
today = datetime.now().date()

# 遍历目录中的所有文件
for filename in os.listdir(log_directory):
    full_path = os.path.join(log_directory, filename)
    
    # 检查是否为文件，如果是文件夹则跳过
    if not os.path.isfile(full_path):
        print(f"跳过目录: {filename}")
        continue

    try:
        file_date_str = ''.join(filter(str.isdigit, filename))[:8]  # 取前 8 位数字作为日期
        file_date = datetime.strptime(file_date_str, '%Y%m%d').date()
    except ValueError:
        print(f"无法解析日期: {filename}, 跳过该文件.")
        continue
    
    # 检查文件日期是否为今天之前
    if file_date < today:
        # 构造目标目录路径
        target_dir = os.path.join(target_root, file_date.strftime('%Y/%m/%d'))
        
        # 如果目标目录不存在，则创建之
        if not os.path.exists(target_dir):
            os.makedirs(target_dir)
        
        # 构造源文件路径和目标文件路径
        source_file = os.path.join(log_directory, filename)
        destination_file = os.path.join(target_dir, filename)
        
        # 移动文件
        shutil.move(source_file, destination_file)
