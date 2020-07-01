import os
import utils
import argparse
import yaml


parser = argparse.ArgumentParser()
parser.add_argument('--service', help='Service name to use', type=str, default=None)
parser.add_argument('--namespace', help='Namespace to use', type=str, default=None)
args = parser.parse_args()


def main():
    src_file = os.path.join(os.getcwd(), "deploy/s3/scality-deployment.yaml")
    dst_file = os.path.join(os.getcwd(), "build/scality-deployment.yaml")
    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = yaml.safe_load(src)
            utils.update_metadata(data, name=args.service, namespace=args.namespace)
            print("Deploying {}".format(dst_file))
            yaml.dump(data, dst, default_flow_style=False)

    utils.apply("deploy/s3/scality-storage.yaml")
    utils.apply(dst_file)


if __name__ == "__main__":
    main()
